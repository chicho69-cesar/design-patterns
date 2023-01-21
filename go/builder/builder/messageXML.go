package builder

import "encoding/xml"

type XMLMessageBuilder struct {
	message Message
}

func (x *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	x.message.Recipient = recipient
	return x
}

func (x *XMLMessageBuilder) SetMessage(message string) MessageBuilder {
	x.message.Text = message
	return x
}

func (x *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(x.message)
	if err != nil {
		return nil, err
	}

	return &MessageRepresented {
		Body: data,
		Format: "XML",
	}, nil
}
