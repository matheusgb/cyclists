package tests

import (
	"testing"

	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/user"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestUpdateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := requests.UpdateUser{
		Name: "Test",
	}

	domain := domains.InitUpdate(
		request.Name,
		"1",
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
}
