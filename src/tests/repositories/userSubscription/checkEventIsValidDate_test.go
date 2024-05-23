package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func TestCheckEventIsValidDateRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs(domain.BikeEventID).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "start_date_registration", "end_date_registration"},
			).AddRow(
				1,
				time.Now().AddDate(0, 0, -1),
				time.Now().AddDate(0, 0, 1),
			),
		)

	repository := repositories.Init(db)
	err := repository.CheckEventIsValidDate(*domain)

	assert.NoError(t, err)
}

func TestCheckEventIsInvalidDateRepositorySuccess(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs(domain.BikeEventID).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "start_date_registration", "end_date_registration"},
			).AddRow(
				1,
				time.Now().AddDate(0, 0, -2),
				time.Now().AddDate(0, 0, -1),
			),
		)

	repository := repositories.Init(db)
	err := repository.CheckEventIsValidDate(*domain)

	assert.Error(t, err)
	assert.Equal(t, "the event subscription is out of date", err.Error())
}

func TestCheckEventIsValidDateRepositoryError(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()

	mock.ExpectQuery("SELECT").
		WithArgs(domain.BikeEventID).
		WillReturnError(db.Error)

	repository := repositories.Init(db)
	err := repository.CheckEventIsValidDate(*domain)

	assert.Error(t, err)
}
