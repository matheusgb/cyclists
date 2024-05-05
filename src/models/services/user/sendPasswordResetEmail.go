package services

import (
	"fmt"

	"github.com/matheusgb/cyclists/src/config"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) SendPasswordResetEmail(domain domains.User) error {
	var email domains.Email

	findedUser, err := user.repository.FindUserByEmail(domain)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	token, err := domains.CreateJWTToken(findedUser)
	if err != nil {
		return fmt.Errorf("error to generate token")
	}

	email.FromName = "no-reply"
	email.FromAddress = config.InitializedConfigs.SendGrid.FromEmail
	email.Subject = "Password Reset"
	email.ToName = findedUser.Name
	email.ToAddress = findedUser.Email
	email.PlainTextContent = fmt.Sprintf("{%s}?email=%s&token=%s", config.InitializedConfigs.SendGrid.URLFrontEnd, email.ToAddress, token)
	email.HtmlContent = fmt.Sprintf("{%s}?email=%s&token=%s", config.InitializedConfigs.SendGrid.URLFrontEnd, email.ToAddress, token)

	err = domains.SendPasswordResetEmail(email)
	if err != nil {
		return err
	}

	return nil
}
