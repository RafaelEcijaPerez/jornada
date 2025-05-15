package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jornada-backend/models"
	"net/http"
	"strings"
	"time"
)

// JornadaService representa el servicio que gestiona las jornadas laborales
type JornadaService struct {
	BaseURL string
	Client  *http.Client
}

// Constructor para JornadaService
func NewJornadaService() *JornadaService {
	return &JornadaService{
		BaseURL: "http://localhost:8080/dolibarr/api/index.php/",
		Client:  &http.Client{},
	}
}

// Extraer token del header Authorization
func (js *JornadaService) GetTokenFromHeader(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("authorization header faltante")
	}
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("formato de token inválido")
	}
	return parts[1], nil
}

// Inicia una jornada laboral
func (js *JornadaService) StartWorkSession(token string, fechaInicio string) (models.ApiResponse, error) {
	startTime, err := time.Parse("2006-01-02T15:04:05Z", fechaInicio)
	if err != nil {
		return models.ApiResponse{}, fmt.Errorf("error al parsear fechaInicio: %v", err)
	}
	fechaInicioFormatted := startTime.UTC().Format("2006-01-02T15:04:05Z")

	payload := map[string]string{
		"fecha_inicio": fechaInicioFormatted,
		"token":        token,
	}

	return js.doPostRequest("jornadasapi/jornadas/start", token, payload)
}

// Finaliza una jornada laboral
func (js *JornadaService) EndWorkSession(token string, fechaFin string) (models.ApiResponse, error) {
	endTime, err := time.Parse(time.RFC3339Nano, fechaFin)
	if err != nil {
		return models.ApiResponse{}, fmt.Errorf("error al parsear fechaFin: %v", err)
	}

	fechaFinFormatted := endTime.UTC().Format("2006-01-02T15:04:05")

	payload := map[string]string{
		"fecha_fin": fechaFinFormatted,
		"token":     token,
	}

	return js.doPostRequest("jornadasapi/jornadas/end", token, payload)
}

// Actualiza el estado de una jornada
func (js *JornadaService) UpdateWorkSession(token string, estado string) (models.ApiResponse, error) {
	payload := map[string]string{
		"estado": estado,
		"token":  token,
	}
	return js.doPostRequest("jornadasapi/jornadas/update", token, payload)
}

// Método reutilizable para hacer POST a la API
func (js *JornadaService) doPostRequest(endpoint string, token string, payload map[string]string) (models.ApiResponse, error) {
	url := js.BaseURL + endpoint

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return models.ApiResponse{}, fmt.Errorf("error al serializar el JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return models.ApiResponse{}, fmt.Errorf("error al crear la solicitud POST: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DOLAPIKEY", token)

	resp, err := js.Client.Do(req)
	if err != nil {
		return models.ApiResponse{}, fmt.Errorf("error al hacer la solicitud a Dolibarr: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ApiResponse{}, fmt.Errorf("error al leer la respuesta de Dolibarr: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return models.ApiResponse{}, fmt.Errorf("error de Dolibarr: %s - %s", resp.Status, string(body))
	}

	var apiResp models.ApiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return models.ApiResponse{}, fmt.Errorf("error al parsear respuesta JSON: %v", err)
	}

	return apiResp, nil
}

// Obtener jornada activa
func (js *JornadaService) GetActiveSession(token string) (models.WorkSession, error) {
	url := js.BaseURL + "jornadasapi/jornadas/active"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.WorkSession{}, fmt.Errorf("error al crear solicitud: %v", err)
	}
	req.Header.Set("DOLAPIKEY", token)

	resp, err := js.Client.Do(req)
	if err != nil {
		return models.WorkSession{}, fmt.Errorf("error al hacer solicitud: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return models.WorkSession{}, fmt.Errorf("respuesta inválida: %s", string(body))
	}

	var session models.WorkSession
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return models.WorkSession{}, fmt.Errorf("error al parsear respuesta: %v", err)
	}

	return session, nil
}

