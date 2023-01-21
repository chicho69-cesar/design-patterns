package main

import (
	"fmt"

	"gitlab.com/chicho69-cesar/prototype/prototype"
)

func main() {
	color := "rojo"
	phones := map[string]string {
		"personal": "1005286",
		"school": "1099207",
	}

	p1 := prototype.Prototype { 
		Name: "Cesar", 
		Age: 21,
		Friends: []string { "Lizeth", "Alonso", "Daniel", "Sammy" },
		Color: &color,
		Phones: phones,
	}
	p2 := p1.Clone()
	
	p2.Name = "Lizeth"
	p2.Age = 20
	p2.Friends[0], p2.Friends[1] = "Joceline", "Jessica"

	//* Estos cambios solo afectan a p1
	color = "azul"
	phones["school"] = "1234567"

	fmt.Println(p1.String())
	fmt.Println(p2.String())
}
