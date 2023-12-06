package engine

import (
	"fmt"
	"log"
	"search_engine/storage"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	Index = "items"
)

type ElasticSearch struct {
	Es elasticsearch.Client
}

func EngineInit() (ElasticSearch, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return ElasticSearch{}, err
	}
	return ElasticSearch{Es: *es}, nil

}

func (es *ElasticSearch) EngineStart() error {
	fmt.Println("Engine Start")
	err := es.SettingsIndex(Index)
	if err != nil {
		return err
	}

	items, err := storage.ReadFile()
	if err != nil {
		return err
	}

	for _, item := range items {
		err := es.Indexing(Index, item)
		if err != nil {
			log.Fatalf("Error indexing document: %s", err)
		}
	}
	fmt.Println("ENGINE START")
	return nil

}
