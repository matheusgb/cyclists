package requests

type CreateUserSubscription struct {
	BikeEventID uint `json:"bike_event_id"`
	UserID      uint `json:"user_id"`
}
