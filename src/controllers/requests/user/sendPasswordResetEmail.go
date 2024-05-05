package requests

type SendPasswordResetEmail struct {
	Email string `json:"email" validate:"required,email"`
}
