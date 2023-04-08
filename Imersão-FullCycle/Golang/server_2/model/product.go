package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	Id   string `json:"id"`
	Name string `json: "name"`
}

type Products struct { // array de Product
	Products []Product
}

//metodo de products | adicionar um novo produto a lista
func (p *Products) Add(product Product) {
	p.Products = append(p.Products, product) // append(tipo, valores)
}

func NewProduct() *Product { // retorna o ponteiro para o novo produto
	return &Product{
		Id: uuid.NewV4().String(),
	}
}
