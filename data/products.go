package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Destcription string  `json:"description"`
	Price        float32 `json:"price" validate:"gt=0"`
	SKU          string  `json:"sku" validate:"required,sku"`
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

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productsList[pos] = p

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productsList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(w io.Reader) error {
	e := json.NewDecoder(w)
	return e.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
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
