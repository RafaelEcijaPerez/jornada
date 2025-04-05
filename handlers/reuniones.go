package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/services"
	"net/http"
)

func MeetingsHandler(w http.ResponseWriter, r *http.Request) {
    // Supongamos que ya tenemos el userID de la sesión o del login
    userID := 1 // Deberías obtener el userID dinámicamente de la autenticación

    // Obtén las reuniones desde Dolibarr
    meetings, err := services.GetMeetingsForUser(userID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error al obtener reuniones: %v", err), http.StatusInternalServerError)
        return
    }

    // Filtra las reuniones para hoy
    meetingsToday := services.FilterMeetingsByDateAndUser(meetings, userID)

    // Devuelve las reuniones filtradas
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(meetingsToday)
}
