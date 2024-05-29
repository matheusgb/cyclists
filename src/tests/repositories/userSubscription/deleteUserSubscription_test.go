package tests

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initDeleteMockedDomain() *domains.UserSubscription {
	request := requests.UserSubscription{
		UserID:      1,
		BikeEventID: 1,
	}

	domain := domains.Init(request.BikeEventID, request.UserID)
	return domain
}

func TestDeleteUserSubscriptionRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE`)).
		WithArgs(domain.BikeEventID, domain.UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repository := repositories.Init(db)
	_, err := repository.DeleteUserSubscription(*domain)

	assert.NoError(t, err)
}

func TestDeleteUserSubscriptionNotFoundRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE`)).
		WithArgs(1, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repository := repositories.Init(db)
	_, err := repository.DeleteUserSubscription(*domain)

	assert.Error(t, err)
	assert.Equal(t, "user subscription with id 1 and 1 not found", err.Error())
}

func TestDeleteUserSubscriptionRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initDeleteMockedDomain()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE`)).
		WithArgs(domain.BikeEventID, domain.ID).
		WillReturnError(db.Error)
	mock.ExpectRollback()

	repository := repositories.Init(db)
	_, err := repository.DeleteUserSubscription(*domain)

	assert.Error(t, err)
}
