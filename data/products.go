package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() Products{
	return productList
}

func GetProductByID(id int) (*Product, error) {
	p,_,_ := findProduct(id)
	return p, nil
}

func AddProduct(p *Product){
	i := productList[len(productList) -1].ID
	p.ID = i+1
	productList = append(productList, p)
}

func UpdateProduct(id int, p*Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		SKU: "fjd34",
	},
	&Product{
		ID: 3,
		Name: "Cappuccino",
		Description: "Espresso based coffee drink that originated in Italy",
		Price: 2.5,
		SKU: "ccp3",
	},
}
