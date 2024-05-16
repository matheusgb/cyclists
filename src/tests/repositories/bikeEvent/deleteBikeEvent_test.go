package tests

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initDeleteMockedDomain() *domains.BikeEvent {
	domain := domains.InitID("1")
	return domain
}

func TestDeleteBikeEventRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
		WithArgs(sqlmock.AnyArg(), domain.ID, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
		WithArgs(sqlmock.AnyArg(), domain.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repository := repositories.Init(db)
	_, err := repository.DeleteBikeEvent(*domain)

	assert.NoError(t, err)
}

func TestDeleteBikeEventRepositoryAdminSuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
		WithArgs(sqlmock.AnyArg(), domain.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
		WithArgs(sqlmock.AnyArg(), domain.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repository := repositories.Init(db)
	_, err := repository.DeleteBikeEventAdmin(*domain)

	assert.NoError(t, err)
}

func TestDeleteBikeEventRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()

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
}

func TestDeleteBikeEventRepositoryAdminError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()

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
	_, err := repository.DeleteBikeEvent(*domain)

	assert.Error(t, err)
}
