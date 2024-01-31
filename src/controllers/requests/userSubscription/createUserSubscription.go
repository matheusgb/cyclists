package requests

type CreateUserSubscription struct {
	BikeEventID uint `json:"bike_event_id" validate:"required"`
	UserID      uint `json:"user_id" validate:"required"`
}
