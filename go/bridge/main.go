package main

import (
	"gitlab.com/chicho69-cesar/bridge/todos"
	"gitlab.com/chicho69-cesar/bridge/list"
)

func main() {
	myToDos := factoryTodo("any")
	rendering := factoryList("bullet")
	myToDos.SetList(rendering)

	myToDos.Add("Qué estudiar?")
	myToDos.Add("Explicarlo con palabras sencillas")
	myToDos.Add("Hacer con lo que aprendimos")
	myToDos.Add("Revisar y mejorar")
	myToDos.Add("Qué estudiar?")
	
	myToDos.Print()
}

func factoryTodo(s string) todos.ToDo {
	if s == "unique" {
		return todos.NewUnique()
	}

	return todos.NewAny()
}

func factoryList(s string) list.List {
	if s == "plain" {
		return list.NewPlain()
	}

	return list.NewBullet('*')
}
