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

// Antes da troca para bcrypt, o cadastro aceitava senhas de até 255 caracteres
// (validate:"min=6,max=255", commit 225fd6d), mas o bcrypt só aceita 72 bytes.
// Estes testes cobrem quem tem senha longa e ainda está com MD5.
func TestLegacyPasswordsLongerThanBcryptLimit(t *testing.T) {
	config.InitializedConfigs = &config.Config{Jwt: config.Jwt{Secret: "test-secret"}}

	for _, size := range []int{73, 100, 255} {
		long := strings.Repeat("a", size)
		sum := md5.Sum([]byte(long))
		legacyHash := hex.EncodeToString(sum[:])

		t.Run(fmt.Sprintf("LoginMigratesLegacyHash/%dBytes", size), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repository := mocks.NewMockIUser(ctrl)
			service := services.Init(repository)

			domain := domains.InitLogin("test@mail.com", long)
			stored := entities.User{Model: gorm.Model{ID: 7}, Email: "test@mail.com", Password: legacyHash}
			repository.EXPECT().FindUserByEmail(*domain).Return(stored, nil)

			var saved domains.User
			repository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(
				func(d domains.User) (entities.User, error) {
					saved = d
					return stored, nil
				})

			token, _, err := service.LoginUser(*domain)

			assert.NoError(t, err)
			assert.NotEmpty(t, token)
			assert.NotEqual(t, legacyHash, saved.Password, "a senha não pode continuar em MD5")
			ok, needsRehash := domains.CheckPassword(saved.Password, long)
			assert.True(t, ok, "o novo hash precisa conferir com a senha completa")
			assert.False(t, needsRehash)
			ok, _ = domains.CheckPassword(saved.Password, long+"x")
			assert.False(t, ok, "senha diferente depois do byte 72 não pode conferir")
		})
	}

	t.Run("HashWrongSuffixAfterByte72IsRejected", func(t *testing.T) {
		long := strings.Repeat("a", 80)
		hash, err := domains.HashLegacyPassword(long)
		assert.NoError(t, err)
		ok, _ := domains.CheckPassword(hash, strings.Repeat("a", 72)+"zzzzzzzz")
		assert.False(t, ok)
		ok, _ = domains.CheckPassword(hash, long)
		assert.True(t, ok)
	})

	t.Run("NewPasswordsStillRejectMoreThan72Bytes", func(t *testing.T) {
		_, err := domains.HashPassword(strings.Repeat("a", domains.MaxPasswordBytes+1))
		assert.Error(t, err)
	})

	t.Run("UpTo72BytesUsesPlainBcrypt", func(t *testing.T) {
		hash, err := domains.HashLegacyPassword(strings.Repeat("a", 72))
		assert.NoError(t, err)
		assert.True(t, strings.HasPrefix(hash, "$2"))
	})
}
