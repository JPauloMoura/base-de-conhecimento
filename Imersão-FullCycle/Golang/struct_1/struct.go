package main

import (
	"errors"
	"fmt"
)

type Pessoa struct { // seria como uma classe
	Nome  string
	Idade int
}

func verificarIdade(p Pessoa) (string, error) { // retorna uma string e um erro
	if p.Idade < 18 {
		return "", errors.New("menor de idade")
	}

	return "maior de idade", nil
}

func main() {
	p1 := Pessoa{
		Nome:  "joao",
		Idade: 2,
	}

	result, err := verificarIdade(p1)

	if err != nil { // tratamento de erros
		fmt.Println(err.Error())
	} else {

		fmt.Println(result)
	}

}
