package gestorturnos

// Contrato
type Contrato struct {
	Tipo       TipoContrato
	Horas      int
	DiasLibres DiasLibres
}

// Tipo de contrato
type TipoContrato string

const (
	Rotatorio       TipoContrato = "Rotatorio"
	JornadaReducida TipoContrato = "JornadaReducida"
	Sustitucion     TipoContrato = "Sustitucion"
	Extraordinario  TipoContrato = "Extraordinario"
)

// Constructor de Contrato
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
