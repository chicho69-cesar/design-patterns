package main

import (
	"fmt"

	"gitlab.com/chicho69-cesar/adapter/adapter"
)

func main() {
	var t string
	fmt.Print("Escribe el metodo de transporte: ")
	fmt.Scan(&t)

	transport := adapter.Factory(t)
	transport.Move()
}
