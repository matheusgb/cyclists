package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func TestCheckEventIsValidDateRepository(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()
	t.Run("Success", func(t *testing.T) {
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
	})

	t.Run("InvalidDate", func(t *testing.T) {
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
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		err := repository.CheckEventIsValidDate(*domain)

		assert.Error(t, err)
	})
}
