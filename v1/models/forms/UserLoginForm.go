package forms

import (
	"html"
	"strings"
)

type UserLoginForm struct {
	Email    string `validate:"required,lte=64" json:"email"`
	Password string `validate:"required,max=32" json:"password"`
}

func (uf *UserLoginForm) Prepare() {
	uf.Email = html.EscapeString(strings.TrimSpace(uf.Email))
	uf.Password = html.EscapeString(strings.TrimSpace(uf.Password))
}
