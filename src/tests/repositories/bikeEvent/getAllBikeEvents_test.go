package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domains "github.com/matheusgb/cyclists/src/models/domains/pagination"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	"github.com/matheusgb/cyclists/src/tests"
	"github.com/stretchr/testify/assert"
)

func initGetAllBikeEventsPaginationMockedDomain() *domains.Pagination {
	domain := &domains.Pagination{
		Page:  0,
		Limit: 0,
		Sort:  "",
	}

	return domain
}

func TestGetAllBikeEventsRepository(t *testing.T) {
	db, mock := tests.MockDatabase()
	domain := initGetAllBikeEventsPaginationMockedDomain()
	t.Run("SuccessWithoutParam", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1).AddRow(2))

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("test").AddRow("test2"))

		repository := repositories.Init(db)
		bikeEvent, err := repository.GetAllBikeEvents(domain, "")

		assert.NoError(t, err)
		assert.Equal(t, 2, len(bikeEvent.Rows.([]entities.BikeEvent)))
		assert.Equal(t, "test", bikeEvent.Rows.([]entities.BikeEvent)[0].Name)
		assert.Equal(t, "test2", bikeEvent.Rows.([]entities.BikeEvent)[1].Name)
		assert.Equal(t, int64(2), bikeEvent.TotalRows)
		assert.Equal(t, 1, bikeEvent.TotalPages)
		assert.Equal(t, 1, bikeEvent.Page)
		assert.Equal(t, 10, bikeEvent.Limit)
		assert.Equal(t, "Id desc", bikeEvent.Sort)
	})

	t.Run("SuccessWithParam", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WithArgs("%test%").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1))

		mock.ExpectQuery("SELECT").
			WithArgs("%test%").
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("test"))

		repository := repositories.Init(db)
		bikeEvent, err := repository.GetAllBikeEvents(domain, "test")

		assert.NoError(t, err)
		assert.Equal(t, 1, len(bikeEvent.Rows.([]entities.BikeEvent)))
		assert.Equal(t, "test", bikeEvent.Rows.([]entities.BikeEvent)[0].Name)
		assert.Equal(t, int64(1), bikeEvent.TotalRows)
		assert.Equal(t, 1, bikeEvent.TotalPages)
		assert.Equal(t, 1, bikeEvent.Page)
		assert.Equal(t, 10, bikeEvent.Limit)
		assert.Equal(t, "Id desc", bikeEvent.Sort)
	})

	t.Run("Error", func(t *testing.T) {

		mock.ExpectQuery("SELECT").
			WithArgs("%test%").
			WillReturnError(db.Error)

		mock.ExpectQuery("SELECT").
			WithArgs("%test%").
			WillReturnError(db.Error)

		repository := repositories.Init(db)
		_, err := repository.GetAllBikeEvents(domain, "test")

		assert.Error(t, err)
	})
}
