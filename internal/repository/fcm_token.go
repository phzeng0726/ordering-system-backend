package repository

import (
	"context"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type FCMTokensRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewFCMTokensRepo(db *gorm.DB, rt *RepoTools) *FCMTokensRepo {
	return &FCMTokensRepo{
		db: db,
		rt: rt,
	}
}

func (r *FCMTokensRepo) Create(ctx context.Context, token domain.FCMToken) error {
	db := r.db.WithContext(ctx)

	if err := db.Create(&token).Error; err != nil {
		return err
	}

	return nil
}

func (r *FCMTokensRepo) GetByUserId(ctx context.Context, userId string) (string, error) {
	var fcmToken domain.FCMToken
	db := r.db.WithContext(ctx)

	if err := db.Where("user_id = ?", userId).First(&fcmToken).Error; err != nil {
		return fcmToken.Token, nil
	}

	return fcmToken.Token, nil
}
