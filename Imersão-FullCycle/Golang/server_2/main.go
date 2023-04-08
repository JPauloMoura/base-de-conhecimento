package main

import (
	"exemplo.com/sever_2/model"
	"exemplo.com/sever_2/server_http"
	uuid "github.com/satori/go.uuid"
)

func main() {
	produto1 := model.Product{
		Id:   uuid.NewV4().String(),
		Name: "Carrinho",
	}

	produto2 := model.Product{
		Id:   uuid.NewV4().String(),
		Name: "Boneca",
	}

	listaProducts := model.Products{}
	listaProducts.Add(produto1)
	listaProducts.Add(produto2)

	// for _, v := range listaProducts.Product {
	// 	fmt.Println(v)
	// }

	//instaciando o servidor
	server := server_http.NewWebServer()
	server.Products = &listaProducts
	server.Run()
}
