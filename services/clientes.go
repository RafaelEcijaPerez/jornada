package services

import (
	"encoding/json"
	"fmt"
	"io"
	"jornada-backend/models"
	"net/http"
	"strconv"
)

// Constante de la URL de la API de Dolibarr
const DolibarrAPI = "http://localhost/dolibarr/api/index.php/thirdparties"

// ObtenerClientes hace la solicitud a Dolibarr y devuelve los clientes
// Recibe el token como argumento, para usarlo en la autenticación
func ObtenerClientes(token string) ([]models.Cliente, error) {
	// Crear la solicitud HTTP GET a la API de Dolibarr
	req, err := http.NewRequest("GET", DolibarrAPI, nil)
	// Crear la solicitud HTTP
	if err != nil {
		return nil, fmt.Errorf("error creando la solicitud: %v", err)
	}

	// Establecer los encabezados necesarios
	req.Header.Set("DOLAPIKEY", token) // Usamos el token recibido
	req.Header.Set("Accept", "application/json")

	// Crear el cliente HTTP y hacer la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud a Dolibarr: %v", err)
	}
	// Asegurarse de cerrar el cuerpo de la respuesta después de leerlo
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta de Dolibarr: %v", resp.Status)
	}

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta de Dolibarr: %v", err)
	}

	// Imprimir la respuesta JSON completa para depuración
	fmt.Println(string(body))

	// Deserializar la respuesta JSON en la estructura de clientes
	var clientes []models.Cliente
	if err := json.Unmarshal(body, &clientes); err != nil {
		return nil, fmt.Errorf("error al parsear la respuesta de Dolibarr: %v", err)
	}

	// Agregar latitud y longitud a cada cliente
	for i := range clientes {
		// Verificar si el cliente tiene opciones de latitud y longitud
		if lat, ok := clientes[i].ArrayOptions["options_latitud"]; ok {
			// Convertir la latitud a float64
			if latFloat, err := strconv.ParseFloat(lat, 64); err == nil {
				clientes[i].Latitud = latFloat
			} else {
				return nil, fmt.Errorf("error al convertir latitud a float64: %v", err)
			}
		}
		// Verificar si el cliente tiene opciones de longitud
		if lon, ok := clientes[i].ArrayOptions["options_longitud"]; ok {
			if lonFloat, err := strconv.ParseFloat(lon, 64); err == nil {
				clientes[i].Longitud = lonFloat
			} else {
				return nil, fmt.Errorf("error al convertir longitud a float64: %v", err)
			}
		}
	}
	return clientes, nil
}

func ObtenerClientePorID(id string, token string) (*models.Cliente, error) {
    // Buscar en lista de clientes
    clientes, err := ObtenerClientes(token)
    if err != nil {
        return nil, err
    }

    for _, c := range clientes {
        if c.ID == id {
            return &c, nil
        }
    }

	return nil, fmt.Errorf("cliente con ID %s no encontrado", id)
}

func EliminarCliente(id string, token string) error {
	// Construir la URL con el ID del cliente
	url := fmt.Sprintf("%s/%s", DolibarrAPI, id)

	// Crear la solicitud HTTP DELETE
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error creando la solicitud DELETE: %v", err)
	}

	// Establecer encabezados
	req.Header.Set("DOLAPIKEY", token)
	req.Header.Set("Accept", "application/json")

	// Ejecutar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al realizar la solicitud DELETE: %v", err)
	}
	defer resp.Body.Close()

	// Verificar código de respuesta
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error al eliminar cliente: %s - %s", resp.Status, string(body))
	}

	return nil
}