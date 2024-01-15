package domains

type User struct {
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

// inicializar no controller com os valores da request
func Init(email, password, passwordConfirmation, name string) *User {
	return &User{
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
		Name:                 name,
	}
}
