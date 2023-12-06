package storage

import (
	"encoding/json"
	"log"
	"os"
	"search_engine/model"
)

const path = "fixed_items.json"

func ReadFile() ([]model.Item, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}

	var jsonData []model.Item
	if err := json.Unmarshal(fileData, &jsonData); err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	return jsonData, nil
}

// func (s *Storage) WriteToJson() error {
// 	items, err := s.SelectItems()
// 	if err != nil {
// 		return fmt.Errorf("Error selecting items: %v", err)
// 	}

// 	jsonItems, err := json.MarshalIndent(items, "", "    ")

// 	if err != nil {
// 		return fmt.Errorf("Error marshalling items: %v", err)
// 	}

// 	err = os.WriteFile("items.json", jsonItems, 0644)
// 	if err != nil {
// 		return fmt.Errorf("Error writing to file:", err)
// 	}

// 	return nil
// }
