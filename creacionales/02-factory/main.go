package main

import (
	"fmt"
	"os"

	"github.com/fenriz07/patrones-go/creacionales/02-factory/factory"
)

func main() {
	var t int

	fmt.Print("Digite la conexión deseada: ")

	_, err := fmt.Scan(&t)

	if err != nil {
		fmt.Printf("error al leer la opción %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(1)

	if conn == nil {
		fmt.Println("Motor no valido")
		os.Exit(1)
	}

	err = conn.Connect()

	if err != nil {
		fmt.Printf("No se puedo crear la conexión: %v", err)
		os.Exit(1)
	}

	now, err := conn.GetNow()

	if err != nil {
		fmt.Printf("No se pudo consultar la fecha: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2006-01-02"))
	err = conn.Close()

	if err != nil {
		fmt.Printf("No se pudo cerrar la conexion, %v", err)
	}

}
