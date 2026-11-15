package application

import (
	"email-verification/config"
	"email-verification/domain"
	"fmt"
	"time"
)

type VerificationService struct {
	repo           domain.VerificationRepository
	emailService   domain.EmailService
	codeGen        domain.CodeGenerator
	codeExpiration time.Duration
}

func NewVerificationService(
	repo domain.VerificationRepository,
	emailService domain.EmailService,
	codeGen domain.CodeGenerator,
	config *config.Config,
) *VerificationService {
	return &VerificationService{
		repo:           repo,
		emailService:   emailService,
		codeGen:        codeGen,
		codeExpiration: config.CodeExpiration,
	}
}

func (s *VerificationService) SendVerification(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if pending, err := s.repo.Get(email); err == nil {
		if time.Now().Before(pending.Exp) {
			return fmt.Errorf("verification already pending")
		}
		// The previous code has expired. Drop it, otherwise this address could never
		// get a new code, because the in-memory repository does not evict anything.
		if err := s.repo.Delete(email); err != nil {
			return fmt.Errorf("failed to delete expired code: %w", err)
		}
	}

	code, err := s.codeGen.Generate()
	if err != nil {
		return err
	}

	if err := s.emailService.SendVerificationCode(email, code); err != nil {
		return err
	}

	verification := domain.Verification{
		Code: s.codeGen.Hash(code),
		Exp:  time.Now().Add(s.codeExpiration),
	}

	return s.repo.Store(email, verification)
}

func (s *VerificationService) VerifyCode(email, code string) error {
	if email == "" || code == "" {
		return fmt.Errorf("email and code cannot be empty")
	}

	verification, err := s.repo.Get(email)
	if err != nil {
		return err
	}

	// Check expiry first. Only the correct code used to clear an expired entry, so a user
	// who had lost the code stayed blocked until the server restarted.
	if time.Now().After(verification.Exp) {
		if err := s.repo.Delete(email); err != nil {
			return fmt.Errorf("failed to delete expired code: %w", err)
		}
		return fmt.Errorf("code expired")
	}

	hashedCode := s.codeGen.Hash(code)
	if verification.Code != hashedCode {
		return fmt.Errorf("invalid code")
	}

	return s.repo.Delete(email)
}
