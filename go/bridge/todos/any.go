package todos

import "gitlab.com/chicho69-cesar/bridge/list"

type Any struct {
	rendering list.List
	todos []string
}

func NewAny() *Any {
	return &Any{}
}

func (a *Any) SetList(l list.List) {
	a.rendering = l
}

func (a *Any) Add(name string) {
	a.todos = append(a.todos, name)
}

func (a *Any) Print() {
	a.rendering.Print(a.todos)
}
