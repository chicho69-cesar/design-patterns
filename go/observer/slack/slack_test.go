package slack

import "testing"

func TestSendMessage(t *testing.T) {
	data := "Hola desde test de go"
	sendMessage(data)
}
