package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

type FCMTokensService struct {
	repo repository.FCMTokens
}

func NewFCMTokensService(repo repository.FCMTokens) *FCMTokensService {
	return &FCMTokensService{repo: repo}
}

func (s *FCMTokensService) Create(ctx context.Context, input CreateTokenInput) error {
	fcmToken := domain.FCMToken{
		UserId:      input.UserId,
		DeviceToken: input.DeviceToken,
	}

	if err := s.repo.Create(ctx, fcmToken); err != nil {
		return err
	}

	return nil
}

func (s *FCMTokensService) Delete(ctx context.Context, input DeleteTokenInput) error {
	fcmToken := domain.FCMToken{
		UserId:      input.UserId,
		DeviceToken: input.DeviceToken,
	}

	if err := s.repo.Delete(ctx, fcmToken); err != nil {
		return err
	}

	return nil
}

func (s *FCMTokensService) GetByUserId(ctx context.Context, userId string) (string, error) {
	menu, err := s.repo.GetByUserId(ctx, userId)
	if err != nil {
		return menu, err
	}

	return menu, nil
}
