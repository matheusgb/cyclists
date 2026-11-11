package tests

import (
	"fmt"
	"strings"
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

func TestCreateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := requests.CreateUser{
		Name:                 "Test",
		Email:                "test@mail.com",
		Password:             "123456",
		PasswordConfirmation: "123456",
	}

	domain := domains.InitCreate(
		request.Email,
		request.Password,
		request.Name,
	)
	repository := mocks.NewMockIUser(ctrl)
	service := services.Init(repository)

	t.Run("Success", func(t *testing.T) {
		entity := entities.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "Test",
			Email:    "test@mail.com",
			Password: domain.Password,
		}

		var saved domains.User
		repository.EXPECT().FindUserByEmail(*domain).Return(entity, fmt.Errorf("record not found"))
		repository.EXPECT().CreateUser(gomock.Any()).DoAndReturn(
			func(d domains.User) (entities.User, error) {
				saved = d
				return entity, nil
			})

		entity, err := service.CreateUser(*domain)

		assert.NoError(t, err)
		assert.NotEqual(t, "123456", saved.Password, "senha nao pode ir em texto puro ao repositorio")
		assert.True(t, strings.HasPrefix(saved.Password, "$2"), "repositorio deve receber hash bcrypt")
		ok, _ := domains.CheckPassword(saved.Password, "123456")
		assert.True(t, ok)
		assert.Equal(t, uint(1), entity.ID)
		assert.Equal(t, "Test", entity.Name)
		assert.Equal(t, "test@mail.com", entity.Email)
	})

	t.Run("UserAlreadyExists", func(t *testing.T) {
		entity := entities.User{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "Test",
			Email:    "test@mail.com",
			Password: domain.Password,
		}

		repository.EXPECT().FindUserByEmail(*domain).Return(entity, nil)
		_, err := service.CreateUser(*domain)

		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("user with email %s already exists", domain.Email))
	})
}
