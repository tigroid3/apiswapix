package forms

type UserRefreshTokenForm struct {
	RefreshToken string `validate:"required" json:"refresh_token"`
}
