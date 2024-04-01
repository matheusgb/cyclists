package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) LoginUser(domain domains.User) (string, error) {
	_, err := user.repository.FindUserByEmailAndPassword(domain)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	return "token", nil
}
