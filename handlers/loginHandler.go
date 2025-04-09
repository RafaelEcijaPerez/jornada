package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/models"
	"jornada-backend/services" 
	"log"
	"net/http"
    "errors"
    "strings"
)

// LoginHandler maneja las solicitudes de login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Verificar que el método de la solicitud sea POST
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Leer el cuerpo de la solicitud
    defer r.Body.Close() // Cierra el cuerpo de la solicitud al final

    // Decodificar el cuerpo JSON en la estructura LoginRequest
    var loginData models.LoginRequest
    // Decodificar el cuerpo JSON en la estructura LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
        http.Error(w, "Error al procesar los datos", http.StatusBadRequest)
        return
    }

    // Validar que los campos no estén vacíos
    if loginData.Usuario == "" || loginData.Contrasena == "" {
        http.Error(w, "Usuario o contraseña no proporcionados", http.StatusBadRequest)
        return
    }

    // Llamar al servicio para obtener el token de DOLIBARR
    token, userID, err := services.GetDolibarrToken(loginData.Usuario, loginData.Contrasena)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error al obtener el token: %v", err), http.StatusUnauthorized)
        return
    }
    // Verificar si el token es válido
    if token == "" {
        http.Error(w, "Token de autorización no válido", http.StatusUnauthorized)
        return
    }
    // Crear la respuesta de login
    response := models.LoginResponse{
        DOLAPIKEY: token,
        Usuario:   loginData.Usuario,
        UserID:    fmt.Sprintf("%d", userID),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// ClientesHandler maneja solicitudes para obtener uno o más clientes
func ClientesHandler(w http.ResponseWriter, r *http.Request) {
    // Verificar que el método de la solicitud sea GET
    if r.Method != http.MethodGet {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Obtener el token desde las cabeceras de la solicitud
    token, err := extraerToken(r.Header.Get("Authorization"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    // Verificar si el token está presente
    if token == "" {
        http.Error(w, "Token de autorización no proporcionado", http.StatusUnauthorized)
        return
    }

    // Obtener el ID del cliente desde la URL
    id := r.URL.Query().Get("id")

    // Verificar si se proporcionó un ID
    if id == "" {
        http.Error(w, "ID de cliente no proporcionado", http.StatusBadRequest)
        return
    }

    // Llamada al servicio de obtener clientes, pasando el token
    var clientes []models.Cliente

    // Verificar si se proporcionó un ID
    if id != "" {
        // Llamar al servicio para obtener cliente por ID
        cliente, err := services.ObtenerClientePorID(id, token)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error al obtener el cliente: %v", err), http.StatusInternalServerError)
            return
        }

        // Verificar si el cliente existe
        if cliente == nil {
            http.Error(w, "Cliente no encontrado", http.StatusNotFound)
            return
        }

        // Agregar el cliente a la lista de clientes
        clientes = append(clientes, *cliente)
    }else {
        // Llamar al servicio para obtener todos los clientes
        clientes, err = services.ObtenerClientes(token)
        if err != nil {
            log.Printf("Error al obtener los clientes: %v", err)
            http.Error(w, fmt.Sprintf("Error al obtener los clientes: %v", err), http.StatusInternalServerError)
            return
        }
    }

    // Responder con la lista de clientes
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Codificar la respuesta en JSON
    if err := json.NewEncoder(w).Encode(clientes); err != nil {
        http.Error(w, fmt.Sprintf("Error al codificar la respuesta: %v", err), http.StatusInternalServerError)
    }
}

// extraerToken valida y extrae el token del header Authorization
func extraerToken(authorizationHeader string) (string, error) {
    if strings.HasPrefix(authorizationHeader, "Bearer ") {
        return strings.TrimPrefix(authorizationHeader, "Bearer "), nil
    }
    return "", errors.New("token de autorización incorrecto")
}