package handlers

import (
	"api-viper/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)
	var res map[string]any

	if err != nil {
		res = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error when trying to insert %v", err),
		}
	} else {
		res = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserted successfully! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
