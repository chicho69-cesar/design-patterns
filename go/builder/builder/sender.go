package builder

type Sender struct {
	builder MessageBuilder
}

func (s *Sender) SetBuilder(m MessageBuilder) {
	s.builder = m
}

func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	m, err := s.builder.SetRecipient(recipient).SetMessage(message).Build()
	return m, err
}
