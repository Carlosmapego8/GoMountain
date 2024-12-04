package auxiliar

import (
	"GoMountain/internal/contrato"
	"GoMountain/internal/diaslibres"
	"GoMountain/internal/turno"
)

type Auxiliar struct {
	ID int
	Nombre string
	Apellidos string
	Contrato contrato.Contrato
	Turnos []turno.Turno
	DiasLibres []diaslibres.DiasLibres
}