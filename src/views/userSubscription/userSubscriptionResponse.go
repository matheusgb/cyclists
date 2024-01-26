package views

import (
	"time"

	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

type UserSubscriptionResponse struct {
	ID          uint      `json:"id"`
	BikeEventID uint      `json:"bike_event_id"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func ConvertUserSubscriptionEntityToResponse(userSubscriptionEntity *entities.UserSubscription) *UserSubscriptionResponse {
	return &UserSubscriptionResponse{
		ID:          userSubscriptionEntity.ID,
		BikeEventID: userSubscriptionEntity.BikeEventID,
		UserID:      userSubscriptionEntity.UserID,
		CreatedAt:   userSubscriptionEntity.CreatedAt,
	}
}
