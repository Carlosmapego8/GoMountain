package diaslibres

type DiasLibres map[Motivo]int

type Motivo int

const (
	AsuntosPropios Motivo = iota
	Vacaciones
	Baja
)
	
