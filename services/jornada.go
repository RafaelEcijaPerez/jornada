package services

import (
	"fmt"
	"jornada-backend/models"
	"time"
	"sync"
)

// Estructura para almacenar las sesiones de trabajo en memoria
var workSessions = make(map[uint]*models.WorkSession)
var mu sync.Mutex // Mutex para manejar la concurrencia

// Función para iniciar una nueva sesión de trabajo
func IniciarSesionTrabajo(userID uint) (*models.WorkSession, error) {
	mu.Lock() // Bloquear acceso concurrente
	defer mu.Unlock()

	startTime := time.Now()

	// Crear una nueva sesión de trabajo en memoria
	workSession := &models.WorkSession{
		ID:        uint(len(workSessions) + 1), // ID incremental
		UserID:    userID,
		StartTime: startTime,
		EndTime:   time.Time{}, // El EndTime estará vacío hasta que se finalice
	}

	// Guardar la sesión de trabajo en memoria
	workSessions[workSession.ID] = workSession

	return workSession, nil
}

// Función para finalizar una sesión de trabajo
func FinalizarSesionTrabajo(sessionID uint) (*models.WorkSession, error) {
	mu.Lock() // Bloquear acceso concurrente
	defer mu.Unlock()

	// Buscar la sesión de trabajo por su ID
	workSession, exists := workSessions[sessionID]
	if !exists {
		return nil, fmt.Errorf("sesión de trabajo no encontrada")
	}

	// Marcar la hora de fin de la sesión
	workSession.EndTime = time.Now()

	// Calcular la duración en segundos
	workSession.CalculateDuration()

	// Actualizar la sesión en memoria (simulando la actualización)
	workSessions[workSession.ID] = workSession

	return workSession, nil
}
