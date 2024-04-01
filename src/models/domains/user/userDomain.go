package domains

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID                   string
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

func InitCreate(email, password, name string) *User {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	password = hex.EncodeToString(hash.Sum(nil))

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
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	password = hex.EncodeToString(hash.Sum(nil))

	return &User{
		Email:    email,
		Password: password,
	}
}
