package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func TestCheckUserIsInEventRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs(domain.ID, domain.BikeEventID, domain.UserID).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "bike_event_id", "user_id"},
			),
		)

	repository := repositories.Init(db)
	err := repository.CheckUserIsInEvent(*domain)

	assert.NoError(t, err)
}

func TestUserIsInEventRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs(domain.ID, domain.BikeEventID, domain.UserID).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "bike_event_id", "user_id"},
			).AddRow(
				domain.ID,
				domain.BikeEventID,
				domain.UserID,
			),
		)

	repository := repositories.Init(db)
	err := repository.CheckUserIsInEvent(*domain)

	assert.Error(t, err)
	assert.Equal(t, "user already subscribed to this event", err.Error())
}

func TestCheckUserIsInEventRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs(domain.ID, domain.BikeEventID, domain.UserID).
		WillReturnError(db.Error)

	repository := repositories.Init(db)
	err := repository.CheckUserIsInEvent(*domain)

	assert.Error(t, err)
}
