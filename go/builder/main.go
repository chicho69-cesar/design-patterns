package main

import (
	"fmt"
	"os"

	"gitlab.com/chicho69-cesar/builder/builder"
	"gitlab.com/chicho69-cesar/builder/factory"
)

func main() {
	var t int

	fmt.Print("Escribe el tipo de mensaje que deseas: ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error al leer el tipo de mensaje: %v", err)
		os.Exit(1)
	}

	sender := &builder.Sender {}
	sender.SetBuilder(factory.Factory(t))
	msg, err := sender.BuildMessage("Gopher", "Hello World!!!")
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(msg.Body))
}
