package adapter

import (
	"gitlab.com/chicho69-cesar/adapter/automovil"
	"gitlab.com/chicho69-cesar/adapter/bycicle"
)

func Factory(s string) Adapter {
	switch s {
		case "automovil":
			d := automovil.Automovil {}
			return &AutomovilAdapter {&d}
		case "bicicleta":
			d := bycicle.Bycicle {}
			return &BycicleAdapter {&d}
	}

	return nil
}
