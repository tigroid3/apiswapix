package forms

import (
	"html"
	"strings"
)

type UserRegisterForm struct {
	Name     string `validate:"required,max=255" json:"name"`
	Phone    string `validate:"max=255" json:"phone"`
	Email    string `validate:"required,lte=64,email" json:"email"`
	Password string `validate:"required,max=32,min=6" json:"password"`
}

func (uf *UserRegisterForm) Prepare() {
	uf.Name = html.EscapeString(strings.TrimSpace(uf.Name))
	uf.Email = html.EscapeString(strings.TrimSpace(uf.Email))
	uf.Password = html.EscapeString(strings.TrimSpace(uf.Password))
	uf.Phone = html.EscapeString(strings.TrimSpace(uf.Phone))
}
