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
	http.HandleFunc("/thirdparties", handlers.ClientesHandler) // Ruta para obtener los clientes
	// Ruta para obtener un cliente específico
	http.HandleFunc("/thirdparties/", handlers.ClienteByIDHandler)
	// Ruta para eliminar un cliente

	// Ruta para iniciar sesión de trabajo
	http.HandleFunc("/work-sessions/start", handlers.StartWorkSessionHandler)
	http.HandleFunc("/work-sessions/end", handlers.EndWorkSessionHandler)
	http.HandleFunc("/work-sessions/update", handlers.UpdateWorkSessionHandler)

}
