package bycicle

import "fmt"

type Bycicle struct {
	Mark  string
	Color string
}

func (b *Bycicle) Advance() {
	fmt.Println("Avanzando!")
}
