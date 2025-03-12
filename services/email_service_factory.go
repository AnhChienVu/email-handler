package services

import "fmt"

func EmailServiceFactory(provider string) (EmailService, error) {
	switch provider {
	case "sendgrid":
		return NewSendGridService(), nil
	default:
		return nil, fmt.Errorf("unsupported email provider: %s", provider)
	}
}
