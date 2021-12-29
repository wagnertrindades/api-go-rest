package main

import (
	"fmt"

	"github.com/wagnertrindades/go-rest-api/models"
	"github.com/wagnertrindades/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando Servidor Rest com Go")
	routes.HandleRequest()
}
