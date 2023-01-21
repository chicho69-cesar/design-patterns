package structures

import (
	"strconv"
	"strings"

	"gitlab.com/chicho69-cesar/composite/composite"
)

type Task struct {
	Name string
	Responsable string
	Price int
	SubTasks []composite.Item
}

func (t *Task) Add(i composite.Item) {
	t.SubTasks = append(t.SubTasks, i)
}

func (t *Task) String() string {
	sb := strings.Builder {}
	sb.WriteString("\t|--")
	sb.WriteString(t.Name)
	sb.WriteString(" - ")
	sb.WriteString(t.Responsable)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(t.GetPrice()))
	sb.WriteString("\n")

	for _, v := range t.SubTasks {
		sb.WriteString(v.String())
	}

	return sb.String()
}

func (t *Task) GetPrice() int {
	total := t.Price
	for _, v := range t.SubTasks {
		total += v.GetPrice()
	}
	return total
}
