package main

import (
	"fmt"

	"github.com/wagnertrindades/go-rest-api/database"
	"github.com/wagnertrindades/go-rest-api/models"
	"github.com/wagnertrindades/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando Servidor Rest com Go")
	routes.HandleRequest()
}
