package adapter

import "github.com/fenriz07/patrones-go/estructurales/Adaptador/auto"

type AutomovilAdapter struct {
	Automovil *auto.Automovil
}

func (a *AutomovilAdapter) Mover() {
	a.Automovil.Encender()
	a.Automovil.Acelerar()
}
