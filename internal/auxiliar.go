package gestorturnos

type Auxiliar struct {
	ID             int
	Contrato       Contrato
	Turnos         []Turno
	TurnoPreferido TipoTurno
	DiasBaja       int
}
