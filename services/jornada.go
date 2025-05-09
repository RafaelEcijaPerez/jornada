package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jornada-backend/models"
	"net/http"
)

const baseURL = "http://localhost:8080/dolibarr/api/index.php/jornadasapi/jornadas"

func StartWorkSession(token string, fechaInicio string) (models.ApiResponse, error) {
	url := fmt.Sprintf("%s/start", baseURL)

	requestBody := models.WorkSession{
		Token:       token,
		FechaInicio: fechaInicio,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return models.ApiResponse{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return models.ApiResponse{}, err
	}
	defer resp.Body.Close()

	var response models.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func EndWorkSession(token string, fechaFin string) (models.ApiResponse, error) {
	url := fmt.Sprintf("%s/end", baseURL)

	requestBody := models.WorkSession{
		Token:    token,
		FechaFin: fechaFin,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return models.ApiResponse{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return models.ApiResponse{}, err
	}
	defer resp.Body.Close()

	var response models.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func UpdateWorkSession(token string, estado string) (models.ApiResponse, error) {
	url := fmt.Sprintf("%s/update", baseURL)

	requestBody := models.WorkSession{
		Token:  token,
		Estado: estado,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return models.ApiResponse{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return models.ApiResponse{}, err
	}
	defer resp.Body.Close()

	var response models.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
