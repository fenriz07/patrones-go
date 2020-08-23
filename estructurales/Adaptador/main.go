package main

import (
	"fmt"

	adapter "github.com/fenriz07/patrones-go/estructurales/Adaptador/Adapter"
)

func main() {

	var t string
	fmt.Println("Digite el tipo de transporte")
	fmt.Scan(&t)

	transport := adapter.Factory(t)

	transport.Mover()

}
