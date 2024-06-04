package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func initFindByEmailAndPasswordMockedDomain() *domains.User {
	request := requests.LoginUser{
		Email:    "test@mail.com",
		Password: "123456",
	}

	domain := domains.InitLogin(request.Email, request.Password)
	return domain
}

func TestFindUserByEmailAndPasswordRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := initFindByEmailAndPasswordMockedDomain()
	t.Run("Success", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WithArgs(domain.Email, domain.Password).
			WillReturnRows(sqlmock.NewRows([]string{"email", "password"}).
				AddRow(domain.Email, domain.Password))

		repository := repositories.Init(db)
		user, err := repository.FindUserByEmailAndPassword(*domain)

		assert.NoError(t, err)
		assert.Equal(t, domain.Email, user.Email)
		assert.Equal(t, domain.Password, user.Password)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.Email, domain.Password).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		_, err := repository.FindUserByEmailAndPassword(*domain)

		assert.Error(t, err)
	})
}
