package tests

import (
	"testing"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/user"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	domain := domains.InitID("1")
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

		repository.EXPECT().GetUser(*domain).Return(entity, nil)
		entity, err := service.GetUser(*domain)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), entity.ID)
		assert.Equal(t, "Test", entity.Name)
		assert.Equal(t, "test@mail.com", entity.Email)
	})
}
