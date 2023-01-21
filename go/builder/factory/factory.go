package factory

import "gitlab.com/chicho69-cesar/builder/builder"

func Factory(t int) builder.MessageBuilder {
	switch t {
		case 1:
			return &builder.JSONMessageBuilder {}
		case 2:
			return &builder.XMLMessageBuilder {}
		default:
			return nil
	}
}
