package validate

import (
	"errors"
	"fmt"
)

var ErrTokenNotValid = errors.New("el usuario no esta logueado")

type ValidateToken struct {
	token string
}

func NewValidatorToken(t string) ValidateToken {
	return ValidateToken{token: t}
}

func (vt *ValidateToken) Validate() error {
	if vt.token != "token-valido" {
		return ErrTokenNotValid
	}

	fmt.Println("Token válido")
	return nil
}
