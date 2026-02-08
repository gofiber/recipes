package email

import (
	"email-verification/config"
	"fmt"
	"net/smtp"
)

type SMTPService struct {
	config *config.Config
}

func NewSMTPService(config *config.Config) *SMTPService {
	return &SMTPService{config: config}
}

func (s *SMTPService) SendVerificationCode(to string, code string) error {
	subject := "Subject: Email Verification Code \n"
	body := fmt.Sprintf("Your verification code is %s", code)
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", s.config.SMTPUser, s.config.SMTPPassword, s.config.SMTPHost)
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", s.config.SMTPHost, s.config.SMTPPort),
		auth,
		s.config.SMTPUser,
		[]string{to},
		message,
	)
	if err != nil {
		return err
	}

	return nil
}
