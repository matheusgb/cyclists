package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func initCreateMockedDomain() *domains.UserSubscription {
	request := requests.UserSubscription{
		UserID:      1,
		BikeEventID: 1,
	}

	domain := domains.Init(request.BikeEventID, request.UserID)
	return domain
}

func TestCreateUserSubscriptionRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := initCreateMockedDomain()
	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("INSERT").
			WithArgs(domain.BikeEventID, domain.UserID, sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		mock.ExpectQuery("SELECT").
			WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id", "bike_event_id", "user_id"},
				).AddRow(
					1,
					domain.BikeEventID,
					domain.UserID,
				),
			)

		repository := repositories.Init(db)
		subscription, err := repository.CreateUserSubscription(*domain)

		assert.NoError(t, err)
		assert.Equal(t, domain.BikeEventID, subscription.BikeEventID)
		assert.Equal(t, domain.UserID, subscription.UserID)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("INSERT").
			WithArgs(
				domain.BikeEventID,
				domain.UserID,
				sqlmock.AnyArg(),
			).
			WillReturnError(db.Error)
		mock.ExpectRollback()

		repository := repositories.Init(db)
		_, err := repository.CreateUserSubscription(*domain)

		assert.Error(t, err)
	})
}
