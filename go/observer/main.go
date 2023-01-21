package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/chicho69-cesar/observer/email"
	"gitlab.com/chicho69-cesar/observer/message"
	"gitlab.com/chicho69-cesar/observer/observer"
	"gitlab.com/chicho69-cesar/observer/slack"
)

func main() {
	m := message.Message{}
	moreObservers := true

	for moreObservers {
		nameObs := readObserver()
		obs := observerFactory(nameObs)
		fmt.Println(obs)
		m.AddObserver(nameObs, obs)

		moreObservers = readAddMore()
	}

	m.Msg = readMessage()
	m.NotifyObservers()
}

func readMessage() string {
	fmt.Print("Digite el mensaje ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("no se pudo leer lo que digitó el usuario: %v", err)
	}

	text = strings.TrimSuffix(text, "\n")
	return text
}

func readObserver() string {
	fmt.Print("Qué observador desea agregar? ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("no se pudo leer lo que digitó el usuario: %v", err)
	}

	text = strings.TrimSuffix(text, "\n")
	return text
}

func readAddMore() bool {
	fmt.Println("Desea agregar más observadores? (s)/(n)")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalf("no se pudo leer lo que digitó el usuario: %v", err)
	}

	if char == 's' {
		return true
	}

	return false
}

func observerFactory(name string) observer.Observer {
	switch name {
		case "slack":
			return &slack.Slack{}
		case "email":
			return &email.Email{}
	}

	return &slack.Slack{}
}
