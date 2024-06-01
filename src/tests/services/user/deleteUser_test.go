package tests

import (
	"testing"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/user"
	"github.com/matheusgb/cyclists/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestDeleteUserService(t *testing.T) {
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

		repository.EXPECT().DeleteUser(*domain).Return(entity, nil)
		_, err := service.DeleteUser(*domain)

		assert.NoError(t, err)
	})
}
