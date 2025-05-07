package models

// LoginRequest representa la estructura del JSON recibido en el login
type LoginRequest struct {
    Login    string `json:"login"`
    Password string `json:"password"`
}


// LoginResponse representa la estructura del JSON de respuesta con el token, el usuario y el ID
type LoginResponse struct {
    DOLAPIKEY string `json:"DOLAPIKEY"`
    Login   string `json:"login"`
    UserID        string `json:"user_id"`  // Agregar el ID del usuario
}
