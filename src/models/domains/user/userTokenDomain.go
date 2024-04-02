package domains

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matheusgb/cyclists/src/config"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func CreateJWTToken(user entities.User) (string, error) {
	secret := config.InitializedConfigs.Jwt.Secret
	thirtyDays := time.Hour * 24 * 30

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(thirtyDays).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return tokenString, nil
}

func VerifyJWTTokenMiddleware(ctx *fiber.Ctx) error {
	secret := config.InitializedConfigs.Jwt.Secret
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(401).JSON(fiber.Map{
			"error": "no token provided",
		})
	}
	token = token[len("Bearer "):]

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	ctx.Locals("user_id", strconv.Itoa(int(claims["id"].(float64))))
	ctx.Locals("user_name", claims["name"])
	ctx.Locals("user_email", claims["email"])
	ctx.Locals("user_role", claims["role"])

	return ctx.Next()
}
