package services

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridService struct{ APIKey string }

func NewSendGridService() *SendGridService {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	return &SendGridService{APIKey: apiKey}
}

func (s *SendGridService) SendEmail(to, subject, body string) error {
	from := mail.NewEmail("Anh", "vuanhchien003@gmail.com")
	toEmail := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, body)
	client := sendgrid.NewSendClient(s.APIKey)

	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Printf("Email sent successfully, Status Code: %d\n", response.StatusCode)
	return nil
}
