package domains

import (
	"fmt"

	"github.com/matheusgb/cyclists/src/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	FromName         string
	FromAddress      string
	ToName           string
	ToAddress        string
	Subject          string
	PlainTextContent string
	HtmlContent      string
}

func SendPasswordResetEmail(email Email) error {
	from := mail.NewEmail(email.FromName, email.FromAddress)
	subject := email.Subject
	to := mail.NewEmail(email.ToName, email.ToAddress)
	plainTextContent := email.PlainTextContent
	htmlContent := email.HtmlContent

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.InitializedConfigs.SendGrid.ApiKey)
	_, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("error to send email: %v", err)
	}
	return nil
}
