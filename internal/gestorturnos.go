package gestorturnos

import (
	"sync"
)

type GestorTurnos struct {
	Auxiliares []Auxiliar
}

// Singleton de GestorTurnos
var instance *GestorTurnos
var once sync.Once

func GetInstance() *GestorTurnos {
	once.Do(func() {
		instance = &GestorTurnos{}
	})
	return instance
}
