package tests

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func initDeleteMockedDomain() *domains.BikeEvent {
	domain := domains.InitID("1")
	return domain
}

func TestDeleteBikeEventRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := initDeleteMockedDomain()
	t.Run("Success", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repository := repositories.Init(db)
		_, err := repository.DeleteBikeEvent(*domain)

		assert.NoError(t, err)
	})

	t.Run("AdminSuccess", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repository := repositories.Init(db)
		_, err := repository.DeleteBikeEventAdmin(*domain)

		assert.NoError(t, err)
	})

	t.Run("NotFound", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), "2").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repository := repositories.Init(db)
		_, err := repository.DeleteBikeEvent(*domain)

		assert.Error(t, err)
		assert.Equal(t, "bike event with id 1 not found", err.Error())
	})

	t.Run("NotFoundAdmin", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), "2").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repository := repositories.Init(db)
		_, err := repository.DeleteBikeEventAdmin(*domain)

		assert.Error(t, err)
		assert.Equal(t, "bike event with id 1 not found", err.Error())
	})

	t.Run("Error", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID, sqlmock.AnyArg()).
			WillReturnError(db.Error)
		mock.ExpectRollback()

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID).
			WillReturnError(db.Error)
		mock.ExpectRollback()

		repository := repositories.Init(db)
		_, err := repository.DeleteBikeEvent(*domain)

		assert.Error(t, err)
	})

	t.Run("AdminError", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID).
			WillReturnError(db.Error)
		mock.ExpectRollback()

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), domain.ID).
			WillReturnError(db.Error)
		mock.ExpectRollback()

		repository := repositories.Init(db)
		_, err := repository.DeleteBikeEventAdmin(*domain)

		assert.Error(t, err)
	})
}
