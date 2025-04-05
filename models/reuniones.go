package models

type Meeting struct {
    ID          int    `json:"id"`
    Date        string `json:"date"`       // La fecha de la reunión, formato "YYYY-MM-DD"
    UserID      int    `json:"user_id"`    // ID del usuario relacionado con la reunión
    Description string `json:"description"` // Descripción de la reunión
}

type GetMeetingsResponse struct {
    Success []Meeting `json:"success"` // Aquí almacenamos las reuniones
    Error   string    `json:"error"`
}
