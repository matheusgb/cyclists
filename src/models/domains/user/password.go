package domains

import (
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// MaxPasswordBytes é o limite do bcrypt: senhas maiores são recusadas em vez
// de truncadas em silêncio.
const MaxPasswordBytes = 72

// HashPassword gera o hash bcrypt (com salt aleatório embutido) da senha.
func HashPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword confere a senha contra o hash armazenado.
//
// Contas criadas antes da troca de algoritmo guardam MD5 hexadecimal sem salt.
// Esse formato só é aceito para conferir o login: quando ele confere, needsRehash
// é true e o chamador deve regravar a senha com HashPassword. O MD5 nunca é usado
// para gerar um novo hash.
func CheckPassword(stored, plain string) (ok bool, needsRehash bool) {
	if strings.HasPrefix(stored, "$2") {
		return bcrypt.CompareHashAndPassword([]byte(stored), []byte(plain)) == nil, false
	}
	if isLegacyMD5(stored) {
		sum := md5.Sum([]byte(plain))
		legacy := hex.EncodeToString(sum[:])
		if subtle.ConstantTimeCompare([]byte(legacy), []byte(stored)) == 1 {
			return true, true
		}
	}
	return false, false
}

func isLegacyMD5(stored string) bool {
	if len(stored) != 32 {
		return false
	}
	_, err := hex.DecodeString(stored)
	return err == nil
}
