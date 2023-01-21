package adapter

import "gitlab.com/chicho69-cesar/adapter/bycicle"

type BycicleAdapter struct {
	Bycicle *bycicle.Bycicle
}

func (b *BycicleAdapter) Move() {
	b.Bycicle.Advance()
}
