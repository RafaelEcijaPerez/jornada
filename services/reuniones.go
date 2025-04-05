package services

import (
	"encoding/json"
	"fmt"
	"io"
	"jornada-backend/models"
	"log"
	"net/http"
	"time"
)


func GetMeetingsForUser(userID int) ([]models.Meeting, error) {
    url := "http://localhost/dolibarr/api/index.php/agendaevents" // Endpoint corregido

    // Crea un cliente HTTP con un timeout
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("error al crear la solicitud HTTP: %v", err)
    }

    // Agregar headers si es necesario (por ejemplo, el token de autenticación)
    req.Header.Set("Authorization", "Bearer your_api_key") // Cambia 'your_api_key' por tu clave real
    req.Header.Set("Accept", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
    }
    defer resp.Body.Close()

    // Leer la respuesta
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error al leer la respuesta: %v", err)
    }

    // Verificar si la respuesta fue exitosa
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error en la respuesta: %s", string(body))
    }

    // Imprimir la respuesta cruda para depuración
    log.Printf("Respuesta cruda: %s", string(body))

    // Decodificar la respuesta JSON
    var result models.GetMeetingsResponse
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
    }

    // Verificar si hubo un error en la respuesta de la API
    if result.Error != "" {
        return nil, fmt.Errorf("error en la API: %s", result.Error)
    }

    return result.Success, nil
}


func FilterMeetingsByDateAndUser(meetings []models.Meeting, userID int) []models.Meeting {
    today := time.Now().Format("2006-01-02") // Formato: "2025-04-04"
    var filteredMeetings []models.Meeting

    for _, meeting := range meetings {
        meetingDate := meeting.Date[:10] // Extraemos solo la parte de la fecha, asumiendo formato "YYYY-MM-DD"
        if meeting.UserID == userID && meetingDate == today {
            filteredMeetings = append(filteredMeetings, meeting)
        }
    }

    return filteredMeetings
}
