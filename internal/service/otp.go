package service

import (
	"context"
	"ordering-system-backend/internal/repository"
)

type OTPService struct {
	repo repository.OTP
}

func NewOTPService(repo repository.OTP) *OTPService {
	return &OTPService{repo: repo}
}

func (s *OTPService) Create(ctx context.Context, input CreateOTPInput) error {
	if err := s.repo.Create(ctx, input.Token, input.Email); err != nil {
		return err
	}
	return nil
}

func (s *OTPService) Verify(ctx context.Context, input VerifyOTPInput) error {
	if err := s.repo.Verify(ctx, input.Token, input.Password); err != nil {
		return err
	}
	return nil
}
