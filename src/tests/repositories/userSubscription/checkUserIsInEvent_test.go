package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCheckUserIsInEventRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := InitCreateMockedDomain()
	t.Run("Success", func(t *testing.T) {

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
	})

	t.Run("UserIsInEvent", func(t *testing.T) {
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
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.ID, domain.BikeEventID, domain.UserID).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		err := repository.CheckUserIsInEvent(*domain)

		assert.Error(t, err)
	})
}
