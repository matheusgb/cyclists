package tests

import (
	"testing"
	"time"

	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/bikeEvent"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetAllBikeEventsService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockIBikeEvent(ctrl)
	service := services.Init(repository)

	t.Run("Success", func(t *testing.T) {
		pagination := &domainsP.Pagination{
			Limit: 1,
			Page:  1,
			Sort:  "",
		}

		response := &domainsP.Pagination{
			Limit:      1,
			Page:       1,
			Sort:       "Id desc",
			TotalRows:  1,
			TotalPages: 1,
			Rows: []entities.BikeEvent{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Name:                  "Test",
					StartPlace:            "Test",
					StartDate:             time.Now(),
					StartDateRegistration: time.Now().AddDate(0, 0, -10),
					EndDateRegistration:   time.Now().AddDate(0, 0, 10),
					Organizer:             1,
				},
			},
		}

		repository.EXPECT().GetAllBikeEvents(pagination, "").Return(response, nil)
		response, err := service.GetAllBikeEvents(pagination, "")

		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.Rows.([]entities.BikeEvent)[0].ID)
		assert.Equal(t, "Test", response.Rows.([]entities.BikeEvent)[0].Name)
		assert.Equal(t, "Test", response.Rows.([]entities.BikeEvent)[0].StartPlace)
		assert.Equal(t, uint(1), response.Rows.([]entities.BikeEvent)[0].Organizer)
		assert.Equal(
			t,
			time.Now().AddDate(0, 0, -10).Format(time.RFC3339),
			response.Rows.([]entities.BikeEvent)[0].StartDateRegistration.Format(time.RFC3339),
		)
		assert.Equal(
			t,
			time.Now().AddDate(0, 0, 10).Format(time.RFC3339),
			response.Rows.([]entities.BikeEvent)[0].EndDateRegistration.Format(time.RFC3339),
		)
	})
}
