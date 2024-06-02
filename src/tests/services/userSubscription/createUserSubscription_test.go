package tests

import (
	"fmt"
	"testing"
	"time"

	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/userSubscription"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserSubscriptionService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := requests.UserSubscription{
		BikeEventID: 1,
		UserID:      1,
	}

	domain := domains.Init(
		request.BikeEventID,
		request.UserID,
	)
	repository := mocks.NewMockIUserSubscription(ctrl)
	service := services.Init(repository)

	t.Run("Success", func(t *testing.T) {
		entity := entities.UserSubscription{
			ID:          1,
			BikeEventID: 1,
			UserID:      1,
			CreatedAt:   time.Now(),
		}

		repository.EXPECT().CheckUserAndEventExists(*domain).Return(nil)
		repository.EXPECT().CheckEventIsValidDate(*domain).Return(nil)
		repository.EXPECT().CheckUserIsInEvent(*domain).Return(nil)
		repository.EXPECT().CheckEventIsFull(*domain).Return(nil)
		repository.EXPECT().CreateUserSubscription(*domain).Return(entity, nil)

		entity, err := service.CreateUserSubscription(*domain)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), entity.ID)
		assert.Equal(t, uint(1), entity.BikeEventID)
		assert.Equal(t, uint(1), entity.UserID)
	})

	t.Run("CheckUserExistsError", func(t *testing.T) {
		repository.EXPECT().CheckUserAndEventExists(*domain).Return(fmt.Errorf("user not found"))

		_, err := service.CreateUserSubscription(*domain)

		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("user not found"))
	})

	t.Run("CheckBikeEventExistsError", func(t *testing.T) {
		repository.EXPECT().CheckUserAndEventExists(*domain).Return(fmt.Errorf("event not found"))

		_, err := service.CreateUserSubscription(*domain)

		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("event not found"))
	})

	t.Run("CheckEventIsValidDateError", func(t *testing.T) {
		repository.EXPECT().CheckUserAndEventExists(*domain).Return(nil)
		repository.EXPECT().CheckEventIsValidDate(*domain).Return(fmt.Errorf("the event subscription is out of date"))

		_, err := service.CreateUserSubscription(*domain)

		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("the event subscription is out of date"))
	})

	t.Run("CheckUserIsInEventError", func(t *testing.T) {
		repository.EXPECT().CheckUserAndEventExists(*domain).Return(nil)
		repository.EXPECT().CheckEventIsValidDate(*domain).Return(nil)
		repository.EXPECT().CheckUserIsInEvent(*domain).Return(fmt.Errorf("user already subscribed to this event"))

		_, err := service.CreateUserSubscription(*domain)

		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("user already subscribed to this event"))
	})

	t.Run("CheckEventIsFullError", func(t *testing.T) {
		repository.EXPECT().CheckUserAndEventExists(*domain).Return(nil)
		repository.EXPECT().CheckEventIsValidDate(*domain).Return(nil)
		repository.EXPECT().CheckUserIsInEvent(*domain).Return(nil)
		repository.EXPECT().CheckEventIsFull(*domain).Return(fmt.Errorf("event is full"))

		_, err := service.CreateUserSubscription(*domain)

		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("event is full"))
	})
}
