package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/models"
	"jornada-backend/services" // Importar el servicio
	"log"
	"net/http"
)

// LoginHandler maneja las solicitudes de login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Acepta solo solicitudes POST
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Verifica el Content-Type de la solicitud
    var loginData models.LoginRequest
    // Decodifica el cuerpo de la solicitud JSON en la estructura LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
        http.Error(w, "Error al procesar los datos", http.StatusBadRequest)
        return
    }

    // Llama a la función para obtener el token Dolibarr, pasando el nombre de usuario y la contraseña
    token, userID, err := services.GetDolibarrToken(loginData.Usuario, loginData.Contrasena)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error al obtener el token: %v", err), http.StatusUnauthorized)
        return
    }

    // Si el token es válido, crea la respuesta con el token y el user_id
    response := models.LoginResponse{
        DOLAPIKEY: token,
        Usuario:   loginData.Usuario,
        UserID:    fmt.Sprintf("%d", userID), // Convierte user_id a string si es necesario
    }

    // Configura la cabecera de respuesta como JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}



func ClientesHandler(w http.ResponseWriter, r *http.Request) {
    // Acepta solo solicitudes GET
    if r.Method != http.MethodGet {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Obtiene el token de la cabecera de la solicitud
    token := r.Header.Get("Authorization")
    if token == "" {
        http.Error(w, "Token de autorización no proporcionado", http.StatusUnauthorized)
        return
    }

    // Verificar si el token tiene el prefijo "Bearer " y extraer solo el token
    if len(token) > 7 && token[:7] == "Bearer " {
        token = token[7:] // Extraer solo el token
    } else {
        http.Error(w, "Token de autorización incorrecto", http.StatusUnauthorized)
        return
    }

    // Obtiene el parámetro "id" de la URL de la consulta, si existe
    id := r.URL.Query().Get("id")

    // Llama a la función para obtener los clientes, pasando el id si está presente
    var clientes []models.Cliente
    var err error
    if id != "" {
        // Si se pasa un id, obtiene ese cliente en particular
        cliente, err := services.ObtenerClientePorID(id, token)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error al obtener el cliente: %v", err), http.StatusInternalServerError)
            return
        }
        clientes = append(clientes, *cliente) // Si es un cliente, lo agregamos al array
    } else {
        // Si no se pasa id, obtiene todos los clientes
        clientes, err = services.ObtenerClientes(token)
    }

    // Si ocurre un error al obtener los clientes, devuelve un error 500
    if err != nil {
        log.Printf("Error al obtener los clientes: %v", err)
        http.Error(w, fmt.Sprintf("Error al obtener los clientes: %v", err), http.StatusInternalServerError)
        return
    }

    // Si no hay clientes, devolver un array vacío
    if len(clientes) == 0 {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode([]models.Cliente{}) // Responder con un array vacío si no hay clientes
        return
    }

    // Configura la cabecera de respuesta como JSON
    w.Header().Set("Content-Type", "application/json")

    // Codifica los clientes en formato JSON y los envía como respuesta
    if err := json.NewEncoder(w).Encode(clientes); err != nil {
        http.Error(w, fmt.Sprintf("Error al codificar la respuesta: %v", err), http.StatusInternalServerError)
    }
}
