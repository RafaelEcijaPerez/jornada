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
    http.HandleFunc("/clients", handlers.ClientesHandler) // Ruta para obtener los clientes
    // Ruta para obtener un cliente específico
    http.HandleFunc("/clients/", handlers.ClienteByIDHandler)

    // Ruta para obtener las reuniones
    http.HandleFunc("/meetings", handlers.MeetingsHandler)

    // Ruta para iniciar sesión de trabajo
    http.HandleFunc("/work-session/start", handlers.IniciarSesionTrabajoHandler) // Ruta para iniciar sesión de trabajo
	// Ruta para finalizar sesión de trabajo
    http.HandleFunc("/work-session/end", handlers.FinalizarSesionTrabajoHandler) // Ruta para finalizar sesión de trabajo
	

}
