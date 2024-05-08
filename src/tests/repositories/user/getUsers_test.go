package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initGetUserMockedDomain() *domains.User {
	domain := domains.InitID("1")

	return domain
}
func TestGetUserRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initGetUserMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).
			AddRow(1))

	repository := repositories.Init(db)
	user, err := repository.GetUser(*domain)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
}

func TestGetUserRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initGetUserMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs("1").
		WillReturnError(db.Error)

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnError(db.Error)

	repository := repositories.Init(db)
	_, err := repository.GetUser(*domain)

	assert.Error(t, err)
}
