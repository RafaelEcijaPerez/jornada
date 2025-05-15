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

// MeetingService representa el servicio que gestiona las reuniones del usuario
type MeetingService struct {
	BaseURL string
	Client  *http.Client
}

// Constructor para MeetingService
func NewMeetingService() *MeetingService {
	return &MeetingService{
		BaseURL: "http://localhost:8080/dolibarr/api/index.php/",
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

// GetMeetingsForUser obtiene todas las reuniones desde la API
func (ms *MeetingService) GetMeetingsForUser(token string) ([]models.Meeting, error) {
	url := ms.BaseURL + "agendaevents"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := ms.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	log.Printf("Respuesta cruda: %s", string(body))

	var result models.GetMeetingsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	if result.Error != "" {
		return nil, fmt.Errorf("error en la API: %s", result.Error)
	}

	return result.Success, nil
}

// FilterMeetingsByDateAndUser filtra las reuniones por fecha actual y usuario
func (ms *MeetingService) FilterMeetingsByDateAndUser(meetings []models.Meeting, userID int) []models.Meeting {
	today := time.Now().Format("2006-01-02")
	var filteredMeetings []models.Meeting

	for _, meeting := range meetings {
		meetingDate := meeting.Date[:10] // Se asume formato "YYYY-MM-DDTHH:mm:ssZ"
		if meeting.UserID == userID && meetingDate == today {
			filteredMeetings = append(filteredMeetings, meeting)
		}
	}

	return filteredMeetings
}
