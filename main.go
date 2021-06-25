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

func (h *productsHandlers) products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

}
func (h *productsHandlers) post(w http.ResponseWriter, r *http.Request) {

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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
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
	http.HandleFunc("/products", productsHandlers.products)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)

	}

}
