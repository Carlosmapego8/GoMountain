package gestorturnos

import (
	"time"
)

type Turno struct {
	Inicio time.Time
	Fin    time.Time
	Tipo   TipoTurno
}

type TipoTurno int

const (
	Manana TipoTurno = iota
	Tarde
	Noche
)
