package adapter

import (
	"github.com/fenriz07/patrones-go/estructurales/Adaptador/auto"
	"github.com/fenriz07/patrones-go/estructurales/Adaptador/bici"
)

func Factory(s string) Adapter {

	switch s {
	case "automovil":
		d := auto.Automovil{}
		return &AutomovilAdapter{
			Automovil: &d,
		}
	case "bicicleta":
		b := bici.Bicicleta{}
		return &BicicletaAdapter{&b}
	}

	return nil
}
