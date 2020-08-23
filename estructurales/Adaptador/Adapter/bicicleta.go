package adapter

import "github.com/fenriz07/patrones-go/estructurales/Adaptador/bici"

type BicicletaAdapter struct {
	Bicicleta *bici.Bicicleta
}

func (b *BicicletaAdapter) Mover() {
	b.Bicicleta.Avanzar()
}
