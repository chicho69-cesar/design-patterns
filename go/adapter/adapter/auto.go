package adapter

import "gitlab.com/chicho69-cesar/adapter/automovil"

type AutomovilAdapter struct {
	Automovil *automovil.Automovil
}

func (a *AutomovilAdapter) Move() {
	a.Automovil.ShutUp()
	a.Automovil.Acelerate()
}
