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

// Constructor
func NewContrato(tipo TipoContrato, horas int, diasLibres diaslibres.DiasLibres) Contrato {

	if horas < 0 {
		horas = 0
	}

	return Contrato{
		Tipo: tipo,
		Horas: horas,
		DiasLibres: diasLibres,
	}
}