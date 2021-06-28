package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Product struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Price        int    `json:"price"`
	Available    bool   `json:"available"`
}
type productsHandlers struct {
	sync.Mutex
	store map[string]Product
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
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var product Product
	err = json.Unmarshal(bodyBytes, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	product.Id = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Lock()
	h.store[product.Name] = product
	defer h.Unlock()

}
func (h *productsHandlers) get(w http.ResponseWriter, r *http.Request) {
	products := make([]Product, len(h.store))
	h.Lock()

	i := 0
	for _, product := range h.store {
		products[i] = product
		i++
	}
	h.Unlock()

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
		store: map[string]Product{
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
