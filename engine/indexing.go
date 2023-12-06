package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"search_engine/model"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (es *ElasticSearch) Indexing(index string, item model.Item) error {
	body, err := json.Marshal(item)
	if err != nil {
		return err
	}

	indexRequest := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(string(body)),
		Refresh: "true",
	}

	resp, err := indexRequest.Do(context.Background(), &(es.Es))
	if err != nil {
		fmt.Printf("Error executing search: %s", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error response: %s", string(body))
	}
	return nil
}

func (es *ElasticSearch) SettingsIndex(index string) error {
	indexSettings := `
	{
		"mappings": {
			"properties": {
				"brand": {
					"type": "text",
					"analyzer": "text_analyzer"
				},
				"category": {
					"type": "text",
					"analyzer": "text_analyzer"
				},
				"full-name": {
					"type": "text",
					"analyzer": "text_analyzer"
				},
				"price": {
					"type": "integer"
				},
				"color": {
					"type": "text",
					"analyzer": "text_analyzer"
				},
				"properties": {
					"type": "text",
					"analyzer": "text_analyzer"
				}
			}
		},
		"settings": {
			"analysis": {
				"analyzer": {
					"text_analyzer": {
						"type": "custom",
						"tokenizer": "standard",
						"filter": ["lowercase", "stemmer_filter", "stopwords_filter", "synonym_filter"]
					}
				},
				"filter": {
					"stemmer_filter": {
						"type": "stemmer",
						"language": "russian"

					},
					"stopwords_filter": {
						"type": "stop",
						"stopwords": ["купить", "заказать", "доставить", "что", "когда", "где", "и"],
						"language": "russian"
					},
					"synonym_filter": {
						"type": "synonym",
						"synonyms_path": "/home/vadim/searchEngine/synonyms.txt"
					}	
				}сдуфк
			}
		}
	}
	`

	indexRequest := esapi.IndicesCreateRequest{
		Index: index,
		Body:  strings.NewReader(indexSettings),
	}

	_, err := indexRequest.Do(context.Background(), es.Es.Transport)
	if err != nil {
		return fmt.Errorf("Error creating index: %s", err)
	}

	return nil
}

func (es *ElasticSearch) IndexExists(index string) (bool, error) {
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), &es.Es)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return false, fmt.Errorf("Error checking index existence: %s", res.String())
	}

	if res.StatusCode == 200 {
		return true, nil
	} else if res.StatusCode == 404 {
		return false, nil
	}

	return false, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
}
