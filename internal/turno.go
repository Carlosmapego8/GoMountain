package gestorturnos

import (
	"time"
)

// Turno
type Turno struct {
	Inicio time.Time
	Fin    time.Time
	Tipo   TipoTurno
}

// Tipo de turno
type TipoTurno string

const (
	Manana TipoTurno = "Mañana"
	Tarde  TipoTurno = "Tarde"
	Noche  TipoTurno = "Noche"
)
