package turno

import (
	"time"
)

type Turno struct {
	Dia time.Time
	HoraInicio string
	HoraFin string
	Tipo TipoTurno
}

type TipoTurno int

const (
    Manana TipoTurno = iota
    Tarde                   
    Noche                    
)