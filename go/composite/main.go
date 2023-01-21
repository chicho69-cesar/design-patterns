package main

import (
	"fmt"

	"gitlab.com/chicho69-cesar/composite/structures"
)

func main() {
	p := structures.Project {Name: "Upload images"}
	t1 := structures.Task {Name: "Mockup", Responsable: "UI Designer", Price: 1000}
	p.Add(&t1)

	t2 := structures.Task {Name: "Markup", Responsable: "Web Designer"}
	st21 := structures.SubTask {Name: "HTML", Price: 300}
	st22 := structures.SubTask {Name: "CSS", Price: 700}
	t2.Add(&st21)
	t2.Add(&st22)

	t3 := structures.Task {Name: "JS", Responsable: "FrontEnd Developer", Price: 1000}

	t4 := structures.Task {Name: "API Backend", Responsable: "Backend Developer"}
	st41 := structures.SubTask {Name: "Authentication", Price: 100}
	st42 := structures.SubTask {Name: "DB connection", Price: 900}
	t4.Add(&st41)
	t4.Add(&st42)

	t5 := structures.Task {Name: "Database", Responsable: "DBA", Price: 1000}

	p.Add(&t1)
	p.Add(&t2)
	p.Add(&t3)
	p.Add(&t4)
	p.Add(&t5)

	fmt.Println(&p)
}
