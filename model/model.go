package model

import "fmt"

type Item struct {
	Brand      string   `json:"brand"`
	Category   string   `json:"category"`
	FullName   string   `json:"full-name"`
	Price      int64    `json:"price"`
	Color      string   `json:"color"`
	Properties []string `json:"properties"`
}

func (it *Item) PrintItem(i int) {
	fmt.Printf("№ %d:\n", i+1)
	fmt.Printf("Брэнд: %s\n", it.Brand)
	fmt.Printf("Категория: %s\n", it.Category)
	fmt.Printf("Полное название: %s\n", it.FullName)
	fmt.Printf("Цена: %d\n", it.Price)
	fmt.Printf("Цвет: %s\n", it.Color)

	fmt.Println("Характеристики:")
	for _, prop := range it.Properties {
		fmt.Printf("  - %s\n", prop)
	}

	fmt.Println()

}
