package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()

	domain := domains.User{
		Name:     "Test",
		Email:    "test@mail.com",
		Password: "123456",
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), domain.Name, domain.Email, domain.Password, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).AddRow(1, domain.Name, domain.Email, domain.Password))

	repository := repositories.Init(db)
	user, err := repository.CreateUser(domain)

	assert.NoError(t, err)
	assert.Equal(t, domain.Name, user.Name)
	assert.Equal(t, domain.Email, user.Email)
	assert.Equal(t, domain.Password, user.Password)
}

func TestCreateUserRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()

	domain := domains.User{
		Name:     "Test",
		Email:    "test@mail.com",
		Password: "123456",
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), domain.Name, domain.Email, domain.Password, sqlmock.AnyArg()).
		WillReturnError(db.Error)
	mock.ExpectRollback()

	repository := repositories.Init(db)
	_, err := repository.CreateUser(domain)

	assert.Error(t, err)
}
