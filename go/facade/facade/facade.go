package facade

import (
	"gitlab.com/chicho69-cesar/facade/email"
	"gitlab.com/chicho69-cesar/facade/storage"
	"gitlab.com/chicho69-cesar/facade/validate"
)

type Facade struct {
	to             			string
	comment        			string
	validatorToken 			validate.ValidateToken
	validatorPermission validate.ValidatorPermission
	store 							storage.Storage
	notificator 				email.Email
}

func New(to, comment, token, user string) Facade {
	return Facade {
		to: to,
		comment: comment,
		validatorToken: validate.NewValidatorToken(token),
		validatorPermission: validate.NewValidatorPermission(user),
		store: storage.NewStorage("mysql"),
		notificator: email.NewEmail(),
	}
}

func (f *Facade) Comment() error {
	if err := f.validatorToken.Validate(); err != nil {
		return err
	}

	if err := f.validatorPermission.Validate(); err != nil {
		return err
	}

	f.store.Save(f.comment)
	f.notificator.Send(f.to, f.comment)

	return nil
}

func (f *Facade) Notify() {
	f.notificator.Send(f.to, f.comment)
}
