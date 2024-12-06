package gestorturnos

// DiasLibres
type DiasLibres map[Motivo]int

// Motivo
type Motivo string

const (
	AsuntosPropios Motivo = "AsuntosPropios"
	Vacaciones     Motivo = "Vacaciones"
)
