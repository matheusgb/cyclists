package domains

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
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

// prehashedPrefix marca hashes de senhas com mais de 72 bytes: bcrypt(base64(sha256(senha))).
// O pré-hash existe porque o bcrypt só lê os 72 primeiros bytes e recusa o resto.
const prehashedPrefix = "bcrypt-sha256$"

// HashLegacyPassword é usado só na migração de contas antigas. O cadastro antigo
// aceitava senhas de até 255 caracteres (commit 225fd6d), então uma conta em MD5
// pode ter senha maior que o limite do bcrypt. Até 72 bytes o resultado é o
// mesmo de HashPassword; acima disso a senha passa por SHA-256 antes do bcrypt,
// para que a senha inteira continue valendo (e não só os 72 primeiros bytes).
func HashLegacyPassword(plain string) (string, error) {
	if len(plain) <= MaxPasswordBytes {
		return HashPassword(plain)
	}
	hash, err := bcrypt.GenerateFromPassword(prehash(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return prehashedPrefix + string(hash), nil
}

func prehash(plain string) []byte {
	sum := sha256.Sum256([]byte(plain))
	return []byte(base64.StdEncoding.EncodeToString(sum[:]))
}

// CheckPassword confere a senha contra o hash armazenado.
//
// Contas criadas antes da troca de algoritmo guardam MD5 hexadecimal sem salt.
// Esse formato só é aceito para conferir o login: quando ele confere, needsRehash
// é true e o chamador deve regravar a senha com HashPassword. O MD5 nunca é usado
// para gerar um novo hash.
func CheckPassword(stored, plain string) (ok bool, needsRehash bool) {
	if strings.HasPrefix(stored, prehashedPrefix) {
		hash := strings.TrimPrefix(stored, prehashedPrefix)
		return bcrypt.CompareHashAndPassword([]byte(hash), prehash(plain)) == nil, false
	}
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
