package todos

import "gitlab.com/chicho69-cesar/bridge/list"

type ToDo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
