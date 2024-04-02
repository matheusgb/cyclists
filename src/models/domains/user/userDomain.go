package domains

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID                   string
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

func InitCreate(email, password, name string) *User {
	return &User{
		Email:    email,
		Password: encryptPassword(password),
		Name:     name,
	}
}

func InitUpdate(name, id string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

func InitID(id string) *User {
	return &User{
		ID: id,
	}
}

func InitLogin(email, password string) *User {
	return &User{
		Email:    email,
		Password: encryptPassword(password),
	}
}

func encryptPassword(password string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// func UserPermissionMiddleware(ctx *fiber.Ctx) error {
// 	contextUserID := ctx.Locals("user_id").(string)
// 	contextUserRole := ctx.Locals("user_role").(string)

// 	UserID := ctx.Params("id", "")
// 	if UserID != contextUserID && contextUserRole != "admin" {
// 		UserID = contextUserID
// 	}
// 	if UserID == "" {
// 		ctx.Status(400).JSON(fiber.Map{
// 			"message": "User ID is required",
// 		})
// 		return nil
// 	}
// 	return ctx.Next()
// }
