package handlers

import (
	"encoding/json"
	"github.com/Schariss/product-api/data"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (p *Products) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func(h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	//h.l.Println("Hello world")
	////data, err := ioutil.ReadAll((*r).Body)
	////same thing as
	//_, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	h.l.Fatal(err)
	//	//rw.WriteHeader(500)
	//	//fmt.Fprintf(rw, "An error occurred")
	//	//rw.Write([]byte("An error occurred"))
	//	//we can replace all that using the built-in method of
	//	http.Error(rw, "An error occurred", http.StatusBadGateway)
	//	return
	//}
	////Send response back
	//io.WriteString(rw, "<h1>Hello from a HandleFunc</h1>\n")
	if r.Method == http.MethodGet {
		h.l.Println("GET", r.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			h.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			h.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			h.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		h.getProductById(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Hello) getProductById(productID int, rw http.ResponseWriter, r *http.Request){
	id := productID
	h.l.Println("[DEBUG] get record id", id)
	prod, err := data.GetProductByID(id)
	err = prod.ToJSON(rw)
	if err != nil {
		h.l.Println("[ERROR] serializing product", err)
	}
}


