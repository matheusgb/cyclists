package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) ResetPassword(domain domains.User) error {
	_, err := user.repository.UpdateUser(domain)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	return nil
}
