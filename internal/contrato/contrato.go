package contrato

import (
	"GoMountain/internal/diaslibres"
)

type Contrato struct {
	Tipo TipoContrato
	Horas int
	DiasLibres diaslibres.DiasLibres
}

type TipoContrato int

const (
	Rotatorio TipoContrato = iota
	JornadaReducida
	Sustitucion
	Extraordinario
)