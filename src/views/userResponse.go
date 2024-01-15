package views

import (
	"time"

	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func userResponse(userEntity *entities.User) *UserResponse {
	return &UserResponse{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
}
