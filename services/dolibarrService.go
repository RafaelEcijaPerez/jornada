package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// GetDolibarrToken realiza la solicitud a la API de Dolibarr
func GetDolibarrToken(username, password string) (string, int, error) {
    // URL de la API de Dolibarr para el login
    url := "http://localhost/dolibarr/api/index.php/login"

    // Crear el cliente HTTP
    loginData := map[string]string{
        "login":    username,
        "password": password,
    }

    // Convertir los datos a JSON
    jsonData, err := json.Marshal(loginData)
    if err != nil {
        return "", 0, fmt.Errorf("error al crear los datos JSON: %v", err)
    }

    // Imprimir los datos JSON para depuración
    log.Printf("Enviando solicitud a Dolibarr con datos: %s", string(jsonData))

    // Crear la solicitud HTTP
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", 0, fmt.Errorf("error al crear la solicitud HTTP: %v", err)
    }

    // Agregar headers necesarios
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", "JornadaBackend/1.0")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", 0, fmt.Errorf("error al hacer la solicitud a Dolibarr: %v", err)
    }
    defer resp.Body.Close()

    log.Printf("Código de respuesta de Dolibarr: %d", resp.StatusCode)

    // Verificar si la respuesta fue exitosa
    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body) // Leer el cuerpo de la respuesta para depuración
        log.Printf("Error en la respuesta de Dolibarr: %s", string(body))
        return "", 0, fmt.Errorf("error en la respuesta de Dolibarr: %s", resp.Status)
    }
    // Leer el cuerpo de la respuesta
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", 0, fmt.Errorf("error al leer la respuesta de Dolibarr: %v", err)
    }

    log.Printf("Respuesta de Dolibarr: %s", string(body))

    // Decodificar la respuesta JSON
    var response map[string]interface{}
    // Imprimir la respuesta cruda para depuración
    if err := json.Unmarshal(body, &response); err != nil {
        return "", 0, fmt.Errorf("error al parsear la respuesta de Dolibarr: %v", err)
    }

    // Verificar si la respuesta contiene el token y user_id
    // Asegúrate de que la respuesta tenga la estructura esperada
    if success, ok := response["success"].(map[string]interface{}); ok {
        if token, ok := success["token"].(string); ok {
            // Asegúrate de extraer también el user_id de manera flexible
            if userID, ok := success["user_id"].(interface{}); ok {
                log.Printf("user_id recibido: %v", userID)
                switch v := userID.(type) {
                case float64:
                    return token, int(v), nil
                case string:
                    // Convierte el user_id a int si es una cadena
                    if userInt, err := strconv.Atoi(v); err == nil {
                        return token, userInt, nil
                    }
                    return "", 0, fmt.Errorf("no se pudo convertir user_id de cadena a entero: %s", v)
                default:
                    return "", 0, fmt.Errorf("user_id tiene un tipo inesperado: %T", v)
                }
            }
        }
    }

    return "", 0, fmt.Errorf("no se encontró el token o user_id en la respuesta")
}
