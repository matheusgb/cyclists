package tests

import (
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

func TestDeleteUserSubscriptionService(t *testing.T) {
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

		repository.EXPECT().DeleteUserSubscription(*domain).Return(entity, nil)
		entity, err := service.DeleteUserSubscription(*domain)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), entity.ID)
		assert.Equal(t, uint(1), entity.BikeEventID)
		assert.Equal(t, uint(1), entity.UserID)
	})

}
