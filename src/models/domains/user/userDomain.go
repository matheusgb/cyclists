package domains

// User é o domínio de usuário. Em InitCreate, InitLogin e InitResetPassword o
// campo Password guarda a senha em texto puro, só em memória; o hash é gerado
// na camada de service (HashPassword) antes de chegar ao repositório.
type User struct {
	ID                   string
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

func InitCreate(email, password, name string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
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

func InitLogin(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func InitSendPasswordResetEmail(email string) *User {
	return &User{
		Email: email,
	}
}

func InitResetPassword(id, password string) *User {
	return &User{
		ID:       id,
		Password: password,
	}
}
