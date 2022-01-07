package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wagnertrindades/go-rest-api/database"
	"github.com/wagnertrindades/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personality

	queryParam := r.URL.Query()

	size, _ := strconv.Atoi(queryParam.Get("size"))
	page, _ := strconv.Atoi(queryParam.Get("page"))

	sort := queryParam.Get("sort")
	direction := queryParam.Get("direction")

	if sort == "" {
		sort = "id"
	}

	if direction == "" {
		direction = "desc"
	}

	database.DB.Limit(size).Offset(page).Order(sort + " " + direction).Find(&personalidades)

	json.NewEncoder(w).Encode(personalidades)
}

func FindPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality

	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)

	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
