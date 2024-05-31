package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func initGetBikeEventMockedDomain() *domains.BikeEvent {
	domain := domains.InitID("1")

	return domain
}
func TestGetBikeEventRepository(t *testing.T) {
	db, mock := mocks.MockDatabase()
	domain := initGetBikeEventMockedDomain()
	t.Run("Success", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WithArgs("1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1))

		mock.ExpectQuery("SELECT").
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"bike_event_id"}).
				AddRow(1))

		repository := repositories.Init(db)
		bikeEvent, err := repository.GetBikeEvent(*domain)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), bikeEvent.ID)
	})

	t.Run("Error", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WithArgs("1").
			WillReturnError(db.Error)

		mock.ExpectQuery("SELECT").
			WithArgs(1).
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		_, err := repository.GetBikeEvent(*domain)

		assert.Error(t, err)
	})
}
