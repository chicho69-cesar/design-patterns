package builder

import (
	"testing"
)

func TestSender_BuildMessage(t *testing.T) {
	sender := &Sender {}
	
	json := &JSONMessageBuilder {}
	xmlf := &XMLMessageBuilder {}

	sender.SetBuilder(json)
	msg, err := sender.BuildMessage("Gopher", "Hello World!!!")
	if err != nil {
		t.Fatalf("BuildMessage(): Unexpected error, we got: %v", err)
	}

	t.Log(string(msg.Body))

	sender.SetBuilder(xmlf)
	msg, err = sender.BuildMessage("Gopher", "Hello World!!!")
	if err != nil {
		t.Fatalf("BuildMessage(): Unexpected error, we got: %v", err)
	}

	t.Log(string(msg.Body))
}
