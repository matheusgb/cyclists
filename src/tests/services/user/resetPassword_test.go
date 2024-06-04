package tests

import (
	"fmt"
	"testing"

	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/user"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestResetPasswordService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := requests.ResetPassword{
		Email:                "test@mail.com",
		Password:             "123456",
		PasswordConfirmation: "123456",
	}

	domain := domains.InitResetPassword(
		"1",
		request.Password,
	)

	repository := mocks.NewMockIUser(ctrl)
	service := services.Init(repository)

	t.Run("Success", func(t *testing.T) {
		entity := entities.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:  "Test",
			Email: "test@mail.com",
		}

		repository.EXPECT().UpdateUser(*domain).Return(entity, nil)

		entity, err := service.UpdateUser(*domain)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), entity.ID)
		assert.Equal(t, "Test", entity.Name)
		assert.Equal(t, "test@mail.com", entity.Email)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		entity := entities.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:  "Test",
			Email: "test@mail.com",
		}

		repository.EXPECT().UpdateUser(*domain).Return(entity, fmt.Errorf("user with id %s not found", domain.ID))

		_, err := service.UpdateUser(*domain)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "user with id 1 not found")
	})
}
