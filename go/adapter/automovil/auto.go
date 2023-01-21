package automovil

import "fmt"

type Automovil struct {
	Mark      string
	Model     uint8
	ShutUpped bool
}

func (a *Automovil) ShutUp() {
	if a.ShutUpped {
		fmt.Println("Ya esta encendido")
		return
	}

	fmt.Println("Encendido!")
}

func (a *Automovil) Acelerate() {
	fmt.Println("Acelerando!")
}
