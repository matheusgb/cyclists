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
		var saved domains.User
		repository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(
			func(d domains.User) (entities.User, error) {
				saved = d
				return entities.User{Model: gorm.Model{ID: 1}}, nil
			})

		err := service.ResetPassword(*domain)

		assert.NoError(t, err)
		assert.Equal(t, "1", saved.ID)
		assert.NotEqual(t, "123456", saved.Password)
		ok, needsRehash := domains.CheckPassword(saved.Password, "123456")
		assert.True(t, ok)
		assert.False(t, needsRehash)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		repository.EXPECT().UpdateUser(gomock.Any()).Return(entities.User{}, fmt.Errorf("user with id %s not found", domain.ID))

		err := service.ResetPassword(*domain)

		assert.EqualError(t, err, "user not found")
	})
}
