package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routerId"`
	ClientID  string     `json:"clienteId"`
	Positions []Position `json:"position"` // lista com todas as posições dessa rota de entrega
}

// a posição da rota é baseada em latitude e logitude
type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

//referente a posição atual do entregador/carro
type PartialRouterPosition struct {
	ID       string    `json:"routerId"`
	ClientID string    `json:"clienteId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"` // quandp for True é porque finalizou a rota
}

func NewRouter() *Route {
	return &Route{}
}

//metodo da struct Route | carregar todas rotas para dentro de Positions[]
func (r *Route) LoadPositions() error {
	if r.ID == "" { // primero verificamos se o ID não veio vazio
		return errors.New("router ID not infomed")
	}

	//abre o arquivo de rotas(lat, log)
	file, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer file.Close() //depois que toda função for executada o arquivo sera fechado -> defer

	lines := bufio.NewScanner(file) // permite fazer a leitura do arquivo
	//é realizado um loop para cada linha do arquivo
	for lines.Scan() {
		position := strings.Split(lines.Text(), ",") //ex: "lat,log" -> ["lat", "long"]

		lat, err := strconv.ParseFloat(position[0], 64)
		if err != nil {
			return nil
		}

		long, err := strconv.ParseFloat(position[1], 64)
		if err != nil {
			return nil
		}

		//adicionando uma nova posição a lista
		r.Positions = append(r.Positions, Position{lat, long}) // append(tipo, valores)
	}
	return nil
}

//retorna um JSON com todas as rotas e informaçoes dela
func (r *Route) ExportJsonPositions() ([]string, error) {
	var partialRouterPosition PartialRouterPosition
	var result []string
	total := len(r.Positions) //quantidade de posições

	//percorre todas as posiçoes e seta seus valores
	for key, value := range r.Positions {
		partialRouterPosition.ID = r.ID
		partialRouterPosition.ClientID = r.ClientID
		partialRouterPosition.Position = []float64{value.Lat, value.Long}
		partialRouterPosition.Finished = false

		if total-1 == key {
			partialRouterPosition.Finished = true
		}

		//converte a struct em um json de bytes
		jsonRouter, err := json.Marshal(partialRouterPosition)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRouter))
	}

	return result, nil
}
