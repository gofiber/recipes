package domain

import "time"

type Verification struct {
	Code string
	Exp  time.Time
}

type VerificationRepository interface {
	Store(email string, verification Verification) error
	Get(email string) (Verification, error)
	Delete(email string) error
}

type EmailService interface {
	SendVerificationCode(to string, code string) error
}
