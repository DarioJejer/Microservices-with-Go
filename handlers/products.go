package handlers

import (
	"app/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to fomat JSON", http.StatusInternalServerError)
	}
}

func (p *Products) PostProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to format JSON", http.StatusBadRequest)
		return
	}

	data.PostProduct(prod)
}
