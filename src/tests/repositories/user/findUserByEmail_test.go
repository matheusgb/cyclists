package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func initFindByEmailMockedDomain() *domains.User {
	request := map[string]string{
		"email": "test@mail.com",
	}

	domain := domains.InitSendPasswordResetEmail(request["email"])
	return domain
}

func TestFindUserByEmailRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := initFindByEmailMockedDomain()
	t.Run("Success", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WithArgs(domain.Email).
			WillReturnRows(sqlmock.NewRows([]string{"email"}).
				AddRow(domain.Email))

		repository := repositories.Init(db)
		user, err := repository.FindUserByEmail(*domain)

		assert.NoError(t, err)
		assert.Equal(t, domain.Email, user.Email)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.Email).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		_, err := repository.FindUserByEmail(*domain)

		assert.Error(t, err)
	})
}
