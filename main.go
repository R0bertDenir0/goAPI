package main

import (
	"encoding/json"
	"net/http"
)

type product struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Price        int    `json:"price"`
	Available    bool   `json:"available"`
}
type productsHandlers struct {
	store map[string]product
}

func (h *productsHandlers) get(w http.ResponseWriter, r *http.Request) {
	products := make([]product, len(h.store))

	i := 0
	for _, product := range h.store {
		products[i] = product
		i++
	}
	jsonBytes, err := json.Marshal(products)
	if err != nil {
		//TODO
	}
	w.Write(jsonBytes)

}
func newProductsHandler() *productsHandlers {
	return &productsHandlers{
		store: map[string]product{
			"ID1": {
				Name:         "product1",
				Manufacturer: "manufacturer1",
				Price:        14,
				Available:    true,
			},
		},
	}
}
func main() {
	productsHandlers := newProductsHandler()
	http.HandleFunc("/products", productsHandlers.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)

	}

}
