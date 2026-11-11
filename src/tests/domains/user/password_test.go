package tests

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"testing"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func legacyMD5(plain string) string {
	sum := md5.Sum([]byte(plain))
	return hex.EncodeToString(sum[:])
}

func TestHashPassword(t *testing.T) {
	t.Run("GeneratesSaltedBcryptHash", func(t *testing.T) {
		first, err := domains.HashPassword("123456")
		assert.NoError(t, err)
		second, err := domains.HashPassword("123456")
		assert.NoError(t, err)

		assert.True(t, strings.HasPrefix(first, "$2"))
		assert.NotEqual(t, "123456", first)
		assert.NotEqual(t, legacyMD5("123456"), first)
		assert.NotEqual(t, first, second, "mesmo input deve gerar hashes diferentes (salt)")
		assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(first), []byte("123456")))
	})

	t.Run("RejectsPasswordLongerThanBcryptLimit", func(t *testing.T) {
		_, err := domains.HashPassword(strings.Repeat("a", domains.MaxPasswordBytes+1))
		assert.Error(t, err)
	})
}

func TestCheckPassword(t *testing.T) {
	bcryptHash, err := domains.HashPassword("123456")
	assert.NoError(t, err)

	t.Run("BcryptCorrectPassword", func(t *testing.T) {
		ok, needsRehash := domains.CheckPassword(bcryptHash, "123456")
		assert.True(t, ok)
		assert.False(t, needsRehash)
	})

	t.Run("BcryptWrongPassword", func(t *testing.T) {
		ok, needsRehash := domains.CheckPassword(bcryptHash, "654321")
		assert.False(t, ok)
		assert.False(t, needsRehash)
	})

	t.Run("LegacyMD5CorrectPasswordAsksForRehash", func(t *testing.T) {
		ok, needsRehash := domains.CheckPassword(legacyMD5("123456"), "123456")
		assert.True(t, ok)
		assert.True(t, needsRehash)
	})

	t.Run("LegacyMD5WrongPassword", func(t *testing.T) {
		ok, needsRehash := domains.CheckPassword(legacyMD5("123456"), "654321")
		assert.False(t, ok)
		assert.False(t, needsRehash)
	})

	t.Run("UnknownHashFormatIsRejected", func(t *testing.T) {
		for _, stored := range []string{"", "123456", "not-a-hash", strings.Repeat("z", 32)} {
			ok, needsRehash := domains.CheckPassword(stored, "123456")
			assert.False(t, ok, stored)
			assert.False(t, needsRehash, stored)
		}
	})
}
