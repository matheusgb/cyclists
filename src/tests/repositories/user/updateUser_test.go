package tests

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initUpdateMockedDomain() *domains.User {
	request := requests.UpdateUser{
		Name: "test",
	}

	domain := domains.InitUpdate(request.Name, "1")
	return domain
}

func TestUpdateUserRepository(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initUpdateMockedDomain()
	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.Name, domain.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
			WithArgs(domain.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, domain.Name))

		repository := repositories.Init(db)
		user, err := repository.UpdateUser(*domain)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), user.ID)
		assert.Equal(t, domain.Name, user.Name)
	})

	t.Run("NotFound", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.Name, 2).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
			WithArgs(domain.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, domain.Name))

		repository := repositories.Init(db)
		_, err := repository.UpdateUser(*domain)

		assert.Error(t, err)
		assert.Equal(t, "user with id 1 not found", err.Error())
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.Name, domain.ID).
			WillReturnError(db.Error)
		mock.ExpectCommit()

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
			WithArgs(domain.ID).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		_, err := repository.UpdateUser(*domain)

		assert.Error(t, err)
	})
}
