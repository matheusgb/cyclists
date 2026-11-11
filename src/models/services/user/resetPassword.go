package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) ResetPassword(domain domains.User) error {
	hash, err := domains.HashPassword(domain.Password)
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	domain.Password = hash

	_, err = user.repository.UpdateUser(domain)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	return nil
}
