package main

import (
	"fmt"
	"os"

	"gitlab.com/chicho69-cesar/factory/factory"
)

func main() {
	var t int
	
	fmt.Print("Escoge la conexion que deseas hacer: ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error al leer la opcion: %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Motor invalido")
		os.Exit(1)
	}

	err = conn.Connect()
	if err != nil {
		fmt.Printf("Error al conectar la base de datos: %v", err)
		os.Exit(1)
	}

	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("No se pudo consultar la fecha: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2023-01-18"))
	err = conn.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexion: %v", err)
	}
}
