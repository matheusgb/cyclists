package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func TestCheckUserAndEventExistsRepository(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := InitCreateMockedDomain()
	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.UserID).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"},
				).AddRow(
					1,
				),
			)

		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"},
				).AddRow(
					1,
				),
			)

		repository := repositories.Init(db)
		err := repository.CheckUserAndEventExists(*domain)

		assert.NoError(t, err)
	})

	t.Run("UserNotExists", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.UserID).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"},
				),
			)

		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"},
				).AddRow(
					1,
				),
			)

		repository := repositories.Init(db)
		err := repository.CheckUserAndEventExists(*domain)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("EventNotExists", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.UserID).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"},
				),
			)

		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"},
				),
			)

		repository := repositories.Init(db)
		err := repository.CheckUserAndEventExists(*domain)

		assert.Error(t, err)
		assert.Equal(t, "event not found", err.Error())
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.UserID).
			WillReturnError(db.Error)

		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		err := repository.CheckEventIsValidDate(*domain)

		assert.Error(t, err)
	})
}
