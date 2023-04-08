package main

import (
	"log"
	"net/http"
)

func main() {
	routes()
	log.Println("Servidor rodando na porta :3001")

	http.ListenAndServe(":3001", nil)
}

func routes() {
	http.HandleFunc("/", wellcome)
}

func wellcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Boa garoto teste!"))
}
