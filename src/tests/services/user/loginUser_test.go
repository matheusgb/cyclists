package tests

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/matheusgb/cyclists/src/config"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	services "github.com/matheusgb/cyclists/src/models/services/user"
	mocks "github.com/matheusgb/cyclists/src/tests/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestLoginUserService(t *testing.T) {
	config.InitializedConfigs = &config.Config{Jwt: config.Jwt{Secret: "test-secret"}}

	bcryptHash, err := domains.HashPassword("123456")
	assert.NoError(t, err)
	sum := md5.Sum([]byte("123456"))
	legacyHash := hex.EncodeToString(sum[:])

	storedWith := func(hash string) entities.User {
		return entities.User{
			Model:    gorm.Model{ID: 7},
			Name:     "Test",
			Email:    "test@mail.com",
			Password: hash,
		}
	}

	t.Run("SuccessWithBcryptDoesNotRewritePassword", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockIUser(ctrl)
		service := services.Init(repository)

		domain := domains.InitLogin("test@mail.com", "123456")
		repository.EXPECT().FindUserByEmail(*domain).Return(storedWith(bcryptHash), nil)

		token, entity, err := service.LoginUser(*domain)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, uint(7), entity.ID)
	})

	t.Run("WrongPasswordIsDenied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockIUser(ctrl)
		service := services.Init(repository)

		domain := domains.InitLogin("test@mail.com", "wrong-password")
		repository.EXPECT().FindUserByEmail(*domain).Return(storedWith(bcryptHash), nil)

		token, _, err := service.LoginUser(*domain)

		assert.Empty(t, token)
		assert.EqualError(t, err, "access denied, check if email and password are correct")
	})

	t.Run("UnknownEmailGetsTheSameError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockIUser(ctrl)
		service := services.Init(repository)

		domain := domains.InitLogin("nobody@mail.com", "123456")
		repository.EXPECT().FindUserByEmail(*domain).Return(entities.User{}, fmt.Errorf("record not found"))

		_, _, err := service.LoginUser(*domain)

		assert.EqualError(t, err, "access denied, check if email and password are correct")
	})

	t.Run("StoredPlainTextIsDenied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockIUser(ctrl)
		service := services.Init(repository)

		domain := domains.InitLogin("test@mail.com", "123456")
		repository.EXPECT().FindUserByEmail(*domain).Return(storedWith("123456"), nil)

		_, _, err := service.LoginUser(*domain)

		assert.Error(t, err)
	})

	t.Run("LegacyMD5LoginUpgradesStoredHashToBcrypt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockIUser(ctrl)
		service := services.Init(repository)

		domain := domains.InitLogin("test@mail.com", "123456")
		repository.EXPECT().FindUserByEmail(*domain).Return(storedWith(legacyHash), nil)

		var saved domains.User
		repository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(
			func(d domains.User) (entities.User, error) {
				saved = d
				return storedWith(d.Password), nil
			})

		token, _, err := service.LoginUser(*domain)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, "7", saved.ID)
		assert.Empty(t, saved.Name, "so a senha deve ser regravada")
		assert.True(t, strings.HasPrefix(saved.Password, "$2"))
		ok, needsRehash := domains.CheckPassword(saved.Password, "123456")
		assert.True(t, ok)
		assert.False(t, needsRehash)
	})

	t.Run("LegacyMD5WrongPasswordDoesNotUpgrade", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockIUser(ctrl)
		service := services.Init(repository)

		domain := domains.InitLogin("test@mail.com", "wrong-password")
		repository.EXPECT().FindUserByEmail(*domain).Return(storedWith(legacyHash), nil)

		_, _, err := service.LoginUser(*domain)

		assert.Error(t, err)
	})

	t.Run("LegacyUpgradeFailureDoesNotBlockLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repository := mocks.NewMockIUser(ctrl)
		service := services.Init(repository)

		domain := domains.InitLogin("test@mail.com", "123456")
		repository.EXPECT().FindUserByEmail(*domain).Return(storedWith(legacyHash), nil)
		repository.EXPECT().UpdateUser(gomock.Any()).Return(entities.User{}, fmt.Errorf("db down"))

		token, _, err := service.LoginUser(*domain)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}
