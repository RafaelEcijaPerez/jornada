package models

type Cliente struct {
    ID           string            `json:"id"`                // Identificador del cliente
    Name         string            `json:"name"`              // Nombre del cliente
    Town         string            `json:"town"`              // Ciudad o pueblo
    Address      string            `json:"address"`           // Dirección
    ArrayOptions map[string]string `json:"array_options"`     // Opciones adicionales (mapa de clave-valor)
    Latitud      float64           `json:"latitud,omitempty"` // Latitud extraída de array_options
    Longitud     float64           `json:"longitud,omitempty"`// Longitud extraída de array_options
}
