package models

// LoginRequest representa la estructura del JSON recibido en el login
type LoginRequest struct {
    Usuario    string `json:"usuario"`
    Contrasena string `json:"contrasena"`
}

// LoginResponse representa la estructura del JSON de respuesta con el token, el usuario y el ID
type LoginResponse struct {
    DOLAPIKEY string `json:"DOLAPIKEY"`
    Usuario   string `json:"usuario"`
    UserID        string `json:"user_id"`  // Agregar el ID del usuario
}
