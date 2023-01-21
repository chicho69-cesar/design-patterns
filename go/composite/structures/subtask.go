package structures

import (
	"fmt"
	"strings"

	"gitlab.com/chicho69-cesar/composite/composite"
)

type SubTask struct {
	Name  string
	Price int
}

func (s *SubTask) Add(i composite.Item) {
	fmt.Println("Ya no acepto más subtareas")
}

func (s *SubTask) String() string {
	sb := strings.Builder {}
	sb.WriteString("\t\t|--")
	sb.WriteString(s.Name)
	sb.WriteString("\n")

	return sb.String()
}

func (s *SubTask) GetPrice() int {
	return s.Price
}
