package tests

import (
	"fmt"
	"testing"
	"time"

	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/bikeEvent"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestCreateBikeEventService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := requests.CreateBikeEvent{
		Name:                  "Test",
		StartPlace:            "Test",
		StartDate:             time.Now().AddDate(0, 0, 20),
		StartDateRegistration: time.Now().AddDate(0, 0, -10),
		EndDateRegistration:   time.Now().AddDate(0, 0, 10),
		Organizer:             1,
	}

	domain := domains.InitCreate(
		request.Name,
		request.StartPlace,
		nil,
		request.StartDate,
		request.StartDateRegistration,
		request.EndDateRegistration,
		request.Organizer,
		nil,
	)
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

		repository.EXPECT().CreateBikeEvent(*domain).Return(entity, nil)

		entity, err := service.CreateBikeEvent(*domain)

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

	t.Run("InvalidStartDate", func(t *testing.T) {
		request := requests.CreateBikeEvent{
			Name:                  "Test",
			StartPlace:            "Test",
			StartDate:             time.Now(),
			StartDateRegistration: time.Now().AddDate(0, 0, -10),
			EndDateRegistration:   time.Now().AddDate(0, 0, 10),
			Organizer:             1,
		}

		domain := domains.InitCreate(
			request.Name,
			request.StartPlace,
			nil,
			request.StartDate,
			request.StartDateRegistration,
			request.EndDateRegistration,
			request.Organizer,
			nil,
		)
		entity, err := service.CreateBikeEvent(*domain)

		assert.Error(t, err)
		assert.Equal(
			t,
			fmt.Errorf("the event start date should be after the end date registration"),
			err,
		)
		assert.Equal(t, entities.BikeEvent{}, entity)
	})

	t.Run("EndDateRegistrationInThePast", func(t *testing.T) {
		request := requests.CreateBikeEvent{
			Name:                  "Test",
			StartPlace:            "Test",
			StartDate:             time.Now(),
			StartDateRegistration: time.Now().AddDate(0, 0, -10),
			EndDateRegistration:   time.Now().AddDate(0, 0, -1),
			Organizer:             1,
		}

		domain := domains.InitCreate(
			request.Name,
			request.StartPlace,
			nil,
			request.StartDate,
			request.StartDateRegistration,
			request.EndDateRegistration,
			request.Organizer,
			nil,
		)
		entity, err := service.CreateBikeEvent(*domain)

		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("the date registration should be in the future"), err)
		assert.Equal(t, entities.BikeEvent{}, entity)
	})

	t.Run("EndDateRegistrationInThePast", func(t *testing.T) {
		request := requests.CreateBikeEvent{
			Name:                  "Test",
			StartPlace:            "Test",
			StartDate:             time.Now().AddDate(0, 0, 20),
			StartDateRegistration: time.Now().AddDate(0, 0, 11),
			EndDateRegistration:   time.Now().AddDate(0, 0, 10),
			Organizer:             1,
		}
		domain := domains.InitCreate(
			request.Name,
			request.StartPlace,
			nil,
			request.StartDate,
			request.StartDateRegistration,
			request.EndDateRegistration,
			request.Organizer,
			nil,
		)
		entity, err := service.CreateBikeEvent(*domain)

		assert.Error(t, err)
		assert.Equal(
			t,
			fmt.Errorf("the start date registration should be before the end date registration"),
			err,
		)
		assert.Equal(t, entities.BikeEvent{}, entity)
	})
}
