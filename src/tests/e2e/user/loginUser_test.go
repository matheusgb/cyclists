package controllers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/config"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	gormdb "github.com/matheusgb/cyclists/src/database"
	"github.com/matheusgb/cyclists/src/layers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func postJSON(t *testing.T, app *fiber.App, path string, body any) (int, map[string]any) {
	t.Helper()
	raw, err := json.Marshal(body)
	require.NoError(t, err)
	req, _ := http.NewRequest("POST", path, bytes.NewReader(raw))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req, -1)
	require.NoError(t, err)
	var out map[string]any
	_ = json.NewDecoder(res.Body).Decode(&out)
	return res.StatusCode, out
}

func md5Hex(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// Login contra Postgres real: cobre o que os testes com repositório mockado não
// mostram, que é a troca do hash MD5 por bcrypt chegar de fato ao banco.
func TestLoginAndLegacyPasswordUpgrade(t *testing.T) {
	app := layers.Setup(config.InitializedConfigs)

	db := gormdb.Init()
	db.Connect(*config.InitializedConfigs)
	client := db.GetClient()

	newUser := func(t *testing.T, label string) (id, email string) {
		t.Helper()
		email = fmt.Sprintf("e2e-%s-%d@example.test", label, time.Now().UnixNano())
		status, body := postJSON(t, app, "/api/v1/register", requests.CreateUser{
			Name: "e2e", Email: email, Password: "123456", PasswordConfirmation: "123456",
		})
		require.EqualValues(t, 201, status)
		id = body["id"].(string)
		t.Cleanup(func() {
			client.Exec("DELETE FROM users WHERE id = ?", id)
		})
		return id, email
	}
	storedPassword := func(t *testing.T, id string) string {
		t.Helper()
		var stored string
		require.NoError(t, client.Raw("SELECT password FROM users WHERE id = ?", id).Scan(&stored).Error)
		return stored
	}
	setPassword := func(t *testing.T, client *gorm.DB, id, hash string) {
		t.Helper()
		require.NoError(t, client.Exec("UPDATE users SET password = ? WHERE id = ?", hash, id).Error)
	}

	t.Run("RegisterStoresBcryptAndLoginReturnsToken", func(t *testing.T) {
		id, email := newUser(t, "bcrypt")
		assert.True(t, strings.HasPrefix(storedPassword(t, id), "$2"))

		status, body := postJSON(t, app, "/api/v1/login", requests.LoginUser{Email: email, Password: "123456"})
		assert.EqualValues(t, 200, status)
		assert.NotEmpty(t, body["token"])
	})

	t.Run("WrongPasswordReturns401", func(t *testing.T) {
		_, email := newUser(t, "wrong")
		status, _ := postJSON(t, app, "/api/v1/login", requests.LoginUser{Email: email, Password: "654321"})
		assert.EqualValues(t, 401, status)
	})

	t.Run("LegacyMD5IsUpgradedInDatabaseOnLogin", func(t *testing.T) {
		id, email := newUser(t, "legacy")
		setPassword(t, client, id, md5Hex("123456"))

		status, _ := postJSON(t, app, "/api/v1/login", requests.LoginUser{Email: email, Password: "123456"})
		assert.EqualValues(t, 200, status)
		assert.True(t, strings.HasPrefix(storedPassword(t, id), "$2"), "hash no banco deve ter virado bcrypt")

		status, _ = postJSON(t, app, "/api/v1/login", requests.LoginUser{Email: email, Password: "123456"})
		assert.EqualValues(t, 200, status, "segundo login usa o hash novo")
	})

	t.Run("LegacyMD5WithLongPasswordIsUpgradedInDatabase", func(t *testing.T) {
		long := strings.Repeat("s", 100)
		id, email := newUser(t, "long")
		setPassword(t, client, id, md5Hex(long))

		status, _ := postJSON(t, app, "/api/v1/login", requests.LoginUser{Email: email, Password: long})
		assert.EqualValues(t, 200, status)
		assert.True(t, strings.HasPrefix(storedPassword(t, id), "bcrypt-sha256$"))

		status, _ = postJSON(t, app, "/api/v1/login", requests.LoginUser{Email: email, Password: long})
		assert.EqualValues(t, 200, status)
		status, _ = postJSON(t, app, "/api/v1/login", requests.LoginUser{Email: email, Password: long + "x"})
		assert.EqualValues(t, 401, status)
	})
}
