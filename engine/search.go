package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"search_engine/model"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (es *ElasticSearch) Search(index, query string) ([]model.Item, error) {
	searchRequest := esapi.SearchRequest{
		Index: []string{index},
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"bool": {
					"must": [
						{
							"multi_match": {
								"query": "%s",
								"fields": ["brand^2", "category^4", "full-name", "color^2", "properties^2"],
								"type": "most_fields",
								"fuzziness": "2"
							}
						}		
					],
					"should": [
						{
						  "bool": {
							"must": [
							  {
								"match_phrase": {
								  "properties": {
									"query": "сезонность зима",
									"boost": 5
								  }
								}
							  }
							]
						  }
						},
						{
						  "bool": {
							"must": [
							  {
								"match_phrase": {
								  "properties": {
									"query": "сезонность лето",
									"boost": 5
								  }
								}
							  }
							]
						  }
						},
						{
						  "bool": {
							"must": [
							  {
								"match_phrase": {
								  "properties": {
									"query": "сезонность весна",
									"boost": 2
								  }
								}
							  }
							]
						  }
						},
						{
							"bool": {
							  "must": [
								{
								  "match_phrase": {
									"properties": {
									  "query": "сезонность осень",
									  "boost": 2
									}
								  }
								}
							  ]
							}
						  },
						{
						  "bool": {
							"must": [
							  {
								"match_phrase": {
								  "properties": {
									"query": "сезонность осень-зима",
									"boost": 1.2
								  }
								}
							  }
							]
						  }
						}

					  ]
				}
			}
		}`, query)),
	}

	res, err := searchRequest.Do(context.Background(), &es.Es)
	if err != nil {
		return nil, fmt.Errorf("Error executing search: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Error searching index: %s", res.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error parsing the search response body: %s", err)
	}

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Error parsing search hits")
	}

	var items []model.Item
	for _, hit := range hits {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Error parsing search hit source")
		}

		properties, ok := source["properties"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("Error parsing 'properties' as []interface{}")
		}

		var propertiesStrings []string
		for _, prop := range properties {
			if propString, ok := prop.(string); ok {
				propertiesStrings = append(propertiesStrings, propString)
			}
		}

		item := model.Item{
			Brand:      source["brand"].(string),
			Category:   source["category"].(string),
			FullName:   source["full-name"].(string),
			Price:      int64(source["price"].(float64)),
			Color:      source["color"].(string),
			Properties: propertiesStrings,
		}
		items = append(items, item)
	}

	return items, nil
}
