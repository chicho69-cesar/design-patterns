package clienttwo

import "gitlab.com/chicho69-cesar/singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
