package main

import (
	"jornada-backend/rutas"
	"log"
	"net/http"
)

func main() {
    // Configurar todas las rutas
    rutas.ConfigurarRutas()
    
    log.Println("Servidor escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
