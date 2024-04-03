package domains

type UserSubscription struct {
	ID          string
	BikeEventID uint
	UserID      uint
}

func Init(bikeEventID, userID uint) *UserSubscription {
	return &UserSubscription{
		BikeEventID: bikeEventID,
		UserID:      userID,
	}
}
