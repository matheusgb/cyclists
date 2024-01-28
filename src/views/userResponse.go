package views

import (
	"time"

	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponseWithBikeEvents struct {
	UserResponse
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	BikeEvents []BikeEventResponse `json:"bike_events"`
}

func ConvertUserEntityToResponseBikeEvent(userEntity *entities.User) *UserResponse {
	return &UserResponse{
		ID:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}

func ConvertUserEntityToResponse(userEntity *entities.User) *UserResponseWithBikeEvents {
	return &UserResponseWithBikeEvents{
		UserResponse: UserResponse{
			ID:    userEntity.ID,
			Name:  userEntity.Name,
			Email: userEntity.Email,
		},
		CreatedAt:  userEntity.CreatedAt,
		UpdatedAt:  userEntity.UpdatedAt,
		BikeEvents: ConvertAllBikeEventsEntityToResponseUser(userEntity.BikeEvents),
	}
}

func ConvertAllUsersEntityToResponseBikeEvent(usersEntity []entities.User) []UserResponse {
	var usersResponse []UserResponse
	for _, userEntity := range usersEntity {
		usersResponse = append(usersResponse, *ConvertUserEntityToResponseBikeEvent(&userEntity))
	}
	return usersResponse
}

func ConvertAllUsersEntityToResponse(usersEntity []entities.User) []UserResponseWithBikeEvents {
	var usersResponse []UserResponseWithBikeEvents
	for _, userEntity := range usersEntity {
		usersResponse = append(usersResponse, *ConvertUserEntityToResponse(&userEntity))
	}
	return usersResponse
}
