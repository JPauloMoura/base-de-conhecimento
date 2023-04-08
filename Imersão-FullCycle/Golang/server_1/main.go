package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("servidor rodando")

	http.HandleFunc("/", HelloHandle)
	http.ListenAndServe(":3001", nil)
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}
