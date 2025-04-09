package handlers

import (
	"encoding/json"
	"fmt"
	"jornada-backend/services"
	"net/http"
	"strings"
)

// ClientesHandler maneja la petición GET /clientes// ClientesHandler maneja las solicitudes para obtener los clientes
func ClientesListHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que el método de la solicitud sea GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener el token desde las cabeceras de la solicitud
	token := r.Header.Get("Authorization")
	// Verificar si el token está presente
	if token == "" {
		http.Error(w, "Token de autorización no proporcionado", http.StatusUnauthorized)
		return
	}

	// Llamada al servicio de obtener clientes, pasando el token
	clientes, err := services.ObtenerClientes(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener los clientes: %v", err), http.StatusInternalServerError)
		return
	}

	// Responder con la lista de clientes
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func ClienteByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token de autorización no proporcionado", http.StatusUnauthorized)
		return
	}

	// Obtener el ID desde la URL
	partes := strings.Split(r.URL.Path, "/")
	if len(partes) < 3 || partes[2] == "" {
		http.Error(w, "ID de cliente no proporcionado", http.StatusBadRequest)
		return
	}
	id := partes[2]

	// Llamar al servicio para obtener cliente
	cliente, err := services.ObtenerClientePorID(id, token)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el cliente: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}