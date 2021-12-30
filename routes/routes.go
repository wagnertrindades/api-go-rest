package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wagnertrindades/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaNovaPersonalidade).Methods("Post")

	log.Fatal(http.ListenAndServe(":8000", r))
}
