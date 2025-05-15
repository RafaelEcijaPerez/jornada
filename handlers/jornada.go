package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/models"
	"jornada-backend/services"
	"net/http"
)

var jornadaService = services.NewJornadaService()

func StartWorkSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	token, err := jornadaService.GetTokenFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var session models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	response, err := jornadaService.StartWorkSession(token, session.FechaInicio)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func EndWorkSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	token, err := jornadaService.GetTokenFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var session models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	response, err := jornadaService.EndWorkSession(token, session.FechaFin)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateWorkSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	token, err := jornadaService.GetTokenFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var session models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Validar el estado de la jornada
	if !isValidEstado(session.Estado) {
		http.Error(w, fmt.Sprintf("Estado inválido: %v", session.Estado), http.StatusBadRequest)
		return
	}

	// Llamar al servicio de actualización
	response, err := jornadaService.UpdateWorkSession(token, session.Estado)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	// Enviar respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Función para validar el estado de la jornada
func isValidEstado(estado string) bool {
	// Los estados válidos
	validStates := []string{"activa", "parada", "finalizada"}

	for _, validState := range validStates {
		if estado == validState {
			return true
		}
	}

	return false
}

func ActiveSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	token, err := jornadaService.GetTokenFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session, err := jornadaService.GetActiveWorkSession(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	if session.ID == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}
