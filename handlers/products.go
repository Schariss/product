package handlers

import (
	"github.com/Schariss/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func(p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func(p *Products) getProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	//1st implementation using Marshal
	//d, err := json.Marshal(lp)
	//2nd implementation using encode from the json library
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	//rw.Write(d)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to decode json", http.StatusInternalServerError)
	}
	//display only values
	//p.l.Printf("prod : %v", prod)
	//display values with fields
	//p.l.Printf("prod : %#v", prod)
	data.AddProduct(prod)
	p.l.Println("updated data : ")
	for i := range data.GetProducts() {
		p.l.Printf("prod : %#v\n--------------------", data.GetProducts()[i])
	}
}

func (p Products) updateProducts(id int, rw http.ResponseWriter, r*http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}