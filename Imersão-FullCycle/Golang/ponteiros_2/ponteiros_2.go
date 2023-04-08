package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) setNome(nome string) {
	p.Nome = nome
	fmt.Println(p.Nome)
}

func (p *Pessoa) setNomePonteiro(nome string) {
	p.Nome = nome
	fmt.Println(p.Nome)
}

func main() {
	pessoa := Pessoa{"jp", 10}

	pessoa.setNome("paulo")  //Pessoa -> p.Nome
	fmt.Println(pessoa.Nome) //"jp"

	pessoa.setNomePonteiro("Maria") //Pessoa -> pessoa.Nome
	fmt.Println(pessoa.Nome)        //"Maria"
}

// "&" mostra o endereço de memoria
// "*" mostra o conteudo daquele endereço de memoria
