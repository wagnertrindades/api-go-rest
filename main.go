package main

import (
	"fmt"

	"github.com/wagnertrindades/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando Servidor Rest com Go")
	routes.HandleRequest()
}
