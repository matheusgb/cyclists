package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) LoginUser(domain domains.User) (string, entities.User, error) {
	findedUser, err := user.repository.FindUserByEmailAndPassword(domain)
	if err != nil {
		return "", findedUser, fmt.Errorf("access denied, check if email and password are correct")
	}

	token, err := domains.CreateJWTToken(findedUser)
	if err != nil {
		return "", findedUser, fmt.Errorf("error to generate access token")
	}

	return token, findedUser, nil
}
