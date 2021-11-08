package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Destcription string  `json:"description"`
	Price        float32 `json:"price"`
	SKU          string  `json:"sku"`
	CreatedOn    string  `json:"-"`
	UpdatedOn    string  `json:"-"`
	DeletedOn    string  `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	return productsList
}

func PostProduct(p *Product) {
	p.ID = nextId()
	p.CreatedOn = time.Now().UTC().String()
	productsList = append(productsList, p)
}

func nextId() int {
	lp := productsList[len(productsList)-1]
	return lp.ID + 1
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(w io.Reader) error {
	e := json.NewDecoder(w)
	return e.Decode(p)
}

var productsList = []*Product{
	&Product{
		ID:           1,
		Name:         "Latte",
		Destcription: "Frothy milky coffe",
		Price:        2.45,
		SKU:          "abc323",
		CreatedOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:           2,
		Name:         "Espresso",
		Destcription: "Short and strong coffe without milk",
		Price:        1.99,
		SKU:          "fdj34",
		CreatedOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
}
