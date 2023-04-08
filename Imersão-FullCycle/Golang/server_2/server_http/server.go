package server_http

import (
	"net/http"

	"exemplo.com/sever_2/model"
	"github.com/labstack/echo/v4" //pacote para subir um servidor web
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

//metodo do WebServer | retorna a lista dos produtos
func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}

//metodo do WebServer | adicionar um novo produto
func (w WebServer) addProduct(c echo.Context) error {

	//criando novo produto
	newProduct := model.NewProduct()

	//junta o body da req com meu newProduct
	err := c.Bind(newProduct)

	if err != nil {
		return err
	}

	w.Products.Add(*newProduct) // adiciona na lista de produtos

	return c.JSON(http.StatusCreated, newProduct)
}

//metodo do WebServer | criar um servidor
func (w WebServer) Run() {
	server := echo.New()
	server.GET("/products", w.getAll)
	server.POST("/products", w.addProduct)
	server.Logger.Fatal(server.Start(":3001"))
}
