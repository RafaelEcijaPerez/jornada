package rutas

import (
    "net/http"
    "jornada-backend/handlers"
)

// ConfigurarRutas define todas las rutas de la aplicación
func ConfigurarRutas() {
    // Ruta para login
    http.HandleFunc("/login", handlers.LoginHandler)  // Ruta para hacer login

    // Ruta para obtener los clientes
    http.HandleFunc("/clientes", handlers.ClientesHandler) // Ruta para obtener los clientes
    // Ruta para obtener un cliente específico
    http.HandleFunc("/clientes/", handlers.ClienteByIDHandler)

    // Ruta para obtener las reuniones
    http.HandleFunc("/meetings", handlers.MeetingsHandler)

}
