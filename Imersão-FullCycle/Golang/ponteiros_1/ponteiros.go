package main

import "fmt"

func main() {
	nome := "joao"

	var endereco_nome *string
	endereco_nome = &nome

	fmt.Println(&endereco_nome) //endereco_nome[ ] -> 0xc000006028
	fmt.Println(endereco_nome)  //endereco_nome[&nome] -> 0xc000006028
	fmt.Println(&nome)          //nome[ ] -> 0xc000006028
	fmt.Println(*endereco_nome) //endereco_nome[&nome] -> 0xc000006028 = "joao"
}

// "&" mostra o endereço de memoria
// "*" mostra o conteudo daquele endereço de memoria
