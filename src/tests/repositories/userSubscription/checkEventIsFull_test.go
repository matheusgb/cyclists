package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func InitCreateMockedDomain() *domains.UserSubscription {
	request := requests.UserSubscription{
		BikeEventID: 1,
		UserID:      1,
	}

	domain := &domains.UserSubscription{
		ID:          "1",
		BikeEventID: request.BikeEventID,
		UserID:      request.UserID,
	}

	return domain
}

func TestCheckEventIsFullRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := InitCreateMockedDomain()
	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnRows(sqlmock.NewRows([]string{"participants_limit"}).
				AddRow(2))

		repository := repositories.Init(db)
		err := repository.CheckEventIsFull(*domain)

		assert.NoError(t, err)
	})

	t.Run("EventIsFull", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnRows(sqlmock.NewRows([]string{"participants_limit"}).
				AddRow(1))

		repository := repositories.Init(db)
		err := repository.CheckEventIsFull(*domain)

		assert.Error(t, err)
		assert.Equal(t, "event is full", err.Error())
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnError(db.Error)

		mock.ExpectQuery("SELECT").
			WithArgs(domain.BikeEventID).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		err := repository.CheckEventIsFull(*domain)

		assert.Error(t, err)
	})
}
