package main

import (
	"fmt"

	"github.com/JPauloMoura/simulator/application/route"
)

func main() {
	myRouter := route.Route{
		ID:       "1",
		ClientID: "01",
	}

	myRouter.LoadPositions()
	for _, position := range myRouter.Positions {
		fmt.Println("-->", position) //--> {-15.8298 -47.92705}
	}

	formattedJSON, _ := myRouter.ExportJsonPositions()
	for _, position := range formattedJSON {
		fmt.Println("-->", position) //--> {"routerId":"1","clienteId":"01","position":[-15.83,-47.92688],"finished":false}
	}

}
