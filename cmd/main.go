package main

import (
	"fmt"
	"log"
	"search_engine/engine"
)

const (
// storagePath = "mongodb://localhost:27017/"
)

func run() error {
	// db, err := storage.ConnectStorage(storagePath)
	// if err != nil {
	// 	return fmt.Errorf("Cnnect db error: %v", err)
	// }
	// defer db.CloseStorage()

	// db.WriteToJson()

	es, err := engine.EngineInit()
	if err != nil {
		return fmt.Errorf("Error init engine: %v", err)
	}
	fmt.Println("Engine already initialized")

	if exists, err := es.IndexExists(engine.Index); err != nil && !exists {

		err = es.EngineStart()
		if err != nil {
			return err
		}

	}
	results, err := es.Search(engine.Index, "черные толстовки от diesel")
	if err != nil {
		log.Fatalf("Error searching: %s", err)
	}

	for i, result := range results {
		if i > 10 {
			break
		}
		result.PrintItem(i)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("error running")
	}
}
