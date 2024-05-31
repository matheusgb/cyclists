package tests

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initDeleteMockedDomain() *domains.User {
	domain := domains.InitID("1")
	return domain
}

func TestDeleteUserRepository(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()
	t.Run("Success", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repository := repositories.Init(db)
		_, err := repository.DeleteUser(*domain)

		assert.NoError(t, err)
	})

	t.Run("NotFound", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), 2).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repository := repositories.Init(db)
		_, err := repository.DeleteUser(*domain)

		assert.Error(t, err)
		assert.Equal(t, "user with id 1 not found", err.Error())
	})

	t.Run("Error", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID).
			WillReturnError(db.Error)
		mock.ExpectRollback()

		repository := repositories.Init(db)
		_, err := repository.DeleteUser(*domain)

		assert.Error(t, err)
	})
}
