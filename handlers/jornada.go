package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/models"
	"jornada-backend/services"
	"net/http"
)

func StartWorkSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var session models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	response, err := services.StartWorkSession(session.Token, session.FechaInicio)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func EndWorkSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var session models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	response, err := services.EndWorkSession(session.Token, session.FechaFin)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateWorkSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var session models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	response, err := services.UpdateWorkSession(session.Token, session.Estado)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}
