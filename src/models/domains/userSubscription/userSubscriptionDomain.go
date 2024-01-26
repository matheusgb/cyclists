package domains

type UserSubscription struct {
	ID          string
	BikeEventID uint
	UserID      uint
}

func InitCreate(bikeEventID, userID uint) *UserSubscription {
	return &UserSubscription{
		BikeEventID: bikeEventID,
		UserID:      userID,
	}
}
