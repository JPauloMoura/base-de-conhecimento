package main

import "fmt"

type arrInt []int

func main() {
	var idades arrInt

	idades = append(idades, 10, 20, 30)

	fmt.Println(idades) //[10 20 30]

	for index, value := range idades {
		fmt.Println(idades[index])
		fmt.Println(value)
	}
}
