package main

import (
	"fmt"

	"github.com/wagnertrindades/go-rest-api/database"
	"github.com/wagnertrindades/go-rest-api/models"
	"github.com/wagnertrindades/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", Bio: "History 1"},
		{Id: 2, Name: "Name 2", Bio: "History 2"},
	}

	database.ConnectWithDatabase()

	fmt.Println("Init server Rest with Go")
	routes.HandleRequest()
}
