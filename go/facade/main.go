package main

import (
	"fmt"
	"log"

	"gitlab.com/chicho69-cesar/facade/facade"
)

func main() {
	showInit()

	token := "token-valido"
	user := "user-blog"
	to := "a@algo.com"
	comment := "muy buen video :)"

	f := facade.New(to, comment, token, user)
	err := f.Comment()
	if err != nil {
		log.Fatal(err)
	}

	f.Notify()

	showFinish()
}

func showInit() {
	fmt.Println(`
		**************************
		* Bienvenido al programa *
		**************************
		`,
	)
}

func showFinish() {
	fmt.Println(`
		**************************
		*  Gracias por utilizar  *
		**************************
		`,
	)
}
