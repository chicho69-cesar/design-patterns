package decorator

import "fmt"

type NewHandlerSayGoodBye struct {}

func NewNewHandlerSayGoodBye() *NewHandlerSayGoodBye {
	return &NewHandlerSayGoodBye {}
}

func (h *NewHandlerSayGoodBye) Process() error {
	fmt.Println("Adios")
	return nil
}
