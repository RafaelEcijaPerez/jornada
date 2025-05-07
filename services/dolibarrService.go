package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetDolibarrToken(login, password string) (string, int, error) {
	log.Printf("Recibiendo login para el usuario: %s", login)

	// Preparar los datos JSON
	data := map[string]string{
		"login":    login,
		"password": password,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", 0, fmt.Errorf("error al convertir los datos a JSON: %v", err)
	}

	// Hacer la solicitud POST a Dolibarr
	resp, err := http.Post("http://localhost:8080/dolibarr/api/index.php/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", 0, fmt.Errorf("error al hacer la solicitud a Dolibarr: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Código de respuesta de Dolibarr: %d", resp.StatusCode)

	// Leer el cuerpo de la respuesta
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	bodyStr := string(bodyBytes)
	log.Printf("Respuesta de Dolibarr: %s", bodyStr)

	// Verificar si es HTML (probable error)
	if len(bodyStr) > 0 && bodyStr[0] == '<' {
		return "", 0, fmt.Errorf("respuesta inesperada en HTML de Dolibarr: %s", bodyStr)
	}

	// Parsear el JSON
	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", 0, fmt.Errorf("error al decodificar JSON: %v", err)
	}

	successData, ok := result["success"].(map[string]interface{})
	if !ok {
		return "", 0, fmt.Errorf("respuesta sin campo 'success': %v", result)
	}

	token, ok := successData["token"].(string)
	if !ok {
		return "", 0, fmt.Errorf("token no encontrado o inválido en la respuesta: %v", successData)
	}

	// Opcional: puedes ignorar el userID o devolver 0 si no lo usas
	userID := 0

	return token, userID, nil
}
