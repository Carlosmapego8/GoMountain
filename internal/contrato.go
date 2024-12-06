package gestorturnos

type Contrato struct {
	Tipo       TipoContrato
	Horas      int
	DiasLibres DiasLibres
}

type TipoContrato int

const (
	Rotatorio TipoContrato = iota
	JornadaReducida
	Sustitucion
	Extraordinario
)

// Constructor
func NewContrato(tipo TipoContrato, horas int, diasLibres DiasLibres) Contrato {

	if horas < 0 {
		horas = 0
	}

	return Contrato{
		Tipo:       tipo,
		Horas:      horas,
		DiasLibres: diasLibres,
	}
}
