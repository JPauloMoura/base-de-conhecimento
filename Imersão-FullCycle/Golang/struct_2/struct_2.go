package main

import (
	"errors"
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

//essa função seria como um metodo da struct Pessoa
func (p Pessoa) verificarIdade() (string, error) {
	if p.Idade < 18 {
		return "", errors.New("menor de idade")
	}
	return "maior de idade", nil
}

func main() {
	p1 := Pessoa{"jp", 10}

	result, err := p1.verificarIdade()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
