package services

import (
	"fmt"
	"log"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) LoginUser(domain domains.User) (string, entities.User, error) {
	errDenied := fmt.Errorf("access denied, check if email and password are correct")

	findedUser, err := user.repository.FindUserByEmail(domain)
	if err != nil {
		return "", findedUser, errDenied
	}

	ok, needsRehash := domains.CheckPassword(findedUser.Password, domain.Password)
	if !ok {
		return "", findedUser, errDenied
	}

	if needsRehash {
		user.upgradeLegacyPassword(findedUser.ID, domain.Password)
	}

	token, err := domains.CreateJWTToken(findedUser)
	if err != nil {
		return "", findedUser, fmt.Errorf("error to generate access token")
	}

	return token, findedUser, nil
}

// upgradeLegacyPassword troca o hash MD5 antigo por bcrypt depois de um login
// válido. Se falhar, o login continua e a troca é tentada no próximo acesso.
func (user *User) upgradeLegacyPassword(id uint, plain string) {
	hash, err := domains.HashPassword(plain)
	if err != nil {
		log.Printf("password rehash skipped for user %d: %v", id, err)
		return
	}
	if _, err := user.repository.UpdateUser(domains.User{ID: fmt.Sprint(id), Password: hash}); err != nil {
		log.Printf("password rehash failed for user %d: %v", id, err)
	}
}
