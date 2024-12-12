package gestorturnos

import "time"

type Auxiliar struct {
	NombreUnico    string
	Contrato       Contrato
	Turnos         map[time.Time]Turno
	TurnoPreferido TipoTurno
	DiasBaja       int
}
