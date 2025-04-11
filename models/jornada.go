package models

import (
	"time"
)

// WorkSession representa una sesión de trabajo
type WorkSession struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"` // Relación con el usuario que tiene la sesión de trabajo
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Duration  int64     `json:"duration"` // Duración en segundos

	// Opcional: Se puede agregar una relación con el usuario si tienes un modelo de usuario
	// User     User      `gorm:"foreignKey:UserID" json:"user"`
}

// Método para calcular la duración en segundos
func (ws *WorkSession) CalculateDuration() {
	ws.Duration = int64(ws.EndTime.Sub(ws.StartTime).Seconds())
}
