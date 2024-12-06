package auxiliar

import (
	"GoMountain/internal/contrato"
	"GoMountain/internal/turno"
)

type Auxiliar struct {
	ID int
	Contrato contrato.Contrato
	Turnos []turno.Turno
	TurnoPreferido turno.TipoTurno
	DiasBaja int
}