package domains

type User struct {
	ID                   string
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

func InitCreate(email, password, passwordConfirmation, name string) *User {
	return &User{
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
		Name:                 name,
	}
}

func InitUpdate(name, id string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

func InitID(id string) *User {
	return &User{
		ID: id,
	}
}
