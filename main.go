package main

import (
	"encoding/json"
	"net/http"
)

type product struct {
	name         string `json:"name"`
	manufacturer string `json:"manufacturer"`
	price        int    `json:"price"`
	available    bool   `json:"available"`
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
				name:         "product1",
				manufacturer: "manufacturer1",
				price:        14,
				available:    true,
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
