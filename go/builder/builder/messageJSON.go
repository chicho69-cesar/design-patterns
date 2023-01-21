package builder

import "encoding/json"

type JSONMessageBuilder struct {
	message Message
}

func (j *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	j.message.Recipient = recipient
	return j
}

func (j *JSONMessageBuilder) SetMessage(message string) MessageBuilder {
	j.message.Text = message
	return j
}

func (j *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(j.message)
	if err != nil {
		return nil, err
	}

	return &MessageRepresented {
		Body: data,
		Format: "JSON",
	}, nil
}
