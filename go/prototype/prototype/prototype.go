package prototype

import (
	"fmt"
)

type Prototype struct {
	Name string
	Age  int
	Friends []string
	Color *string
	Phones map[string]string
}

func (p *Prototype) String() string {
	return fmt.Sprintf(
		"Nombre: %s, Edad: %d, Amigos: %v, Color: %s, Telefonos: %v",
		p.Name,
		p.Age,
		p.Friends,
		*p.Color,
		p.Phones,
	)
}

func (p *Prototype) Clone() Prototype {
	/* friends := make([]string, 0, len(p.Friends))
	friends = append([]string {}, p.Friends...) */

	friends := make([]string, len(p.Friends))
	copy(friends, p.Friends)

	color := *p.Color

	phones := make(map[string]string)
	for k, v := range p.Phones {
		phones[k] = v
	}

	return Prototype {
		Name: p.Name,
		Age: p.Age,
		Friends: friends,
		Color: &color,
		Phones: phones,
	}
}
