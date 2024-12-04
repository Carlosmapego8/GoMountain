package contrato

type Contrato struct {
	Tipo TipoContrato
	Horas int
}

type TipoContrato int

const (
	Rotatorio TipoContrato = iota
	JornadaReducida
	Sustitucion
	Extraordinario
)