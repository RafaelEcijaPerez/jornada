package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/models"
	"jornada-backend/services"
	"net/http"
)

// IniciarSesionTrabajoHandler maneja las solicitudes para iniciar una sesión de trabajo
func IniciarSesionTrabajoHandler(w http.ResponseWriter, r *http.Request) {
	// Acepta solo solicitudes POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el cuerpo de la solicitud (requiere el userID)
	var sessionData models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&sessionData); err != nil {
		http.Error(w, "Error al procesar los datos", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para iniciar la sesión de trabajo
	workSession, err := services.IniciarSesionTrabajo(sessionData.UserID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al iniciar la sesión de trabajo: %v", err), http.StatusInternalServerError)
		return
	}

	// Configurar la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workSession)
}

// FinalizarSesionTrabajoHandler maneja las solicitudes para finalizar una sesión de trabajo
func FinalizarSesionTrabajoHandler(w http.ResponseWriter, r *http.Request) {
	// Acepta solo solicitudes POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el cuerpo de la solicitud (requiere el sessionID)
	var sessionData models.WorkSession
	if err := json.NewDecoder(r.Body).Decode(&sessionData); err != nil {
		http.Error(w, "Error al procesar los datos", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para finalizar la sesión de trabajo
	workSession, err := services.FinalizarSesionTrabajo(sessionData.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al finalizar la sesión de trabajo: %v", err), http.StatusInternalServerError)
		return
	}

	// Configurar la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workSession)
}
