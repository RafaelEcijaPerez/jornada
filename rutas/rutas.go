package rutas

import (
	"jornada-backend/handlers"
	"net/http"
)

// ConfigurarRutas define todas las rutas de la aplicación
func ConfigurarRutas() {
	// Ruta para login
	http.HandleFunc("/login", handlers.LoginHandler) // Ruta para hacer login

	// Ruta para obtener los clientes
	http.HandleFunc("/clients", handlers.ClientesHandler) // Ruta para obtener los clientes
	// Ruta para obtener un cliente específico
	http.HandleFunc("/clients/", handlers.ClienteByIDHandler)
	// Ruta para eliminar un cliente
	// Esto es válido con net/http (sin mux)
	http.HandleFunc("/clientes/", handlers.ClienteDeleteHandler)

	// Ruta para obtener las reuniones
	http.HandleFunc("/meetings", handlers.MeetingsHandler)

	// Ruta para iniciar sesión de trabajo
	http.HandleFunc("/work-session/start", handlers.IniciarSesionTrabajoHandler) // Ruta para iniciar sesión de trabajo
	// Ruta para finalizar sesión de trabajo
	http.HandleFunc("/work-session/end", handlers.FinalizarSesionTrabajoHandler) // Ruta para finalizar sesión de trabajo

}
