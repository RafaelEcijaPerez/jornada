package models

type WorkSession struct {
	ID          int    `json:"id,omitempty"`
	Token       string `json:"token"`
	FechaInicio string `json:"fecha_inicio,omitempty"`
	FechaFin    string `json:"fecha_fin,omitempty"`
	Estado      string `json:"estado,omitempty"`
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
