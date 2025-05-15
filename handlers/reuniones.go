package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/services"
	"net/http"
)

func GetTodayMeetingsHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token requerido", http.StatusUnauthorized)
		return
	}

	userID := 123 // Simulación, en producción extraer del token u otro método

	ms := services.NewMeetingService()

	meetings, err := ms.GetMeetingsForUser(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener reuniones: %v", err), http.StatusInternalServerError)
		return
	}

	filtered := ms.FilterMeetingsByDateAndUser(meetings, userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

