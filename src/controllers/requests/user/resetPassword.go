package requests

type ResetPassword struct {
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6,max=100"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}
