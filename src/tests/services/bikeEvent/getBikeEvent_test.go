package tests

import (
	"testing"
	"time"

	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/bikeEvent"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetBikeEventService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	domain := domains.InitID("1")
	repository := mocks.NewMockIBikeEvent(ctrl)
	service := services.Init(repository)

	t.Run("Success", func(t *testing.T) {
		entity := entities.BikeEvent{
			Model: gorm.Model{
				ID: 1,
			},
			Name:                  "Test",
			StartPlace:            "Test",
			StartDate:             time.Now(),
			StartDateRegistration: time.Now().AddDate(0, 0, -10),
			EndDateRegistration:   time.Now().AddDate(0, 0, 10),
			Organizer:             1,
		}

		repository.EXPECT().GetBikeEvent(*domain).Return(entity, nil)

		entity, err := service.GetBikeEvent(*domain)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), entity.ID)
		assert.Equal(t, "Test", entity.Name)
		assert.Equal(t, "Test", entity.StartPlace)
		assert.Equal(t, uint(1), entity.Organizer)
		assert.Equal(
			t,
			time.Now().AddDate(0, 0, -10).Format(time.RFC3339),
			entity.StartDateRegistration.Format(time.RFC3339),
		)
		assert.Equal(
			t,
			time.Now().AddDate(0, 0, 10).Format(time.RFC3339),
			entity.EndDateRegistration.Format(time.RFC3339),
		)
		assert.Equal(t, time.Now().Format(time.RFC3339), entity.StartDate.Format(time.RFC3339))
	})
}
