package structures

import (
	"strconv"
	"strings"

	"gitlab.com/chicho69-cesar/composite/composite"
)

type Project struct {
	Name  string
	Tasks []composite.Item
}

func (p *Project) Add(i composite.Item) {
	p.Tasks = append(p.Tasks, i)
}

func (p *Project) String() string {
	sb := strings.Builder {}
	sb.WriteString(p.Name)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(p.GetPrice()))
	sb.WriteString("\n")

	for _, v := range p.Tasks {
		sb.WriteString(v.String())
	}

	return sb.String()
}

func (p *Project) GetPrice() int {
	price := 0
	for _, v := range p.Tasks {
		price += v.GetPrice()
	}
	return price
}
