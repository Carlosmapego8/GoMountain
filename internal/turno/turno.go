package turno

type Turno struct {
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