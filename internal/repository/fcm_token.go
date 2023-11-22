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

func (r *FCMTokensRepo) Delete(ctx context.Context, token domain.FCMToken) error {
	db := r.db.WithContext(ctx)

	if err := db.Where("user_id = ? AND token = ?", token.UserId, token.DeviceToken).Delete(&token).Error; err != nil {
		return err
	}

	return nil
}

func (r *FCMTokensRepo) GetByUserId(ctx context.Context, userId string) (string, error) {
	var fcmToken domain.FCMToken
	db := r.db.WithContext(ctx)

	if err := db.Where("user_id = ?", userId).Order("created_at DESC").First(&fcmToken).Error; err != nil {
		return fcmToken.DeviceToken, nil
	}

	return fcmToken.DeviceToken, nil
}

func (r *FCMTokensRepo) GetAllBySeatId(ctx context.Context, seatId int) ([]string, error) {
	var fcmTokens []domain.FCMToken
	var deviceTokens []string
	db := r.db.WithContext(ctx)

	sqlQuery := "SELECT ft.*" +
		" FROM fcm_tokens ft" +
		" INNER JOIN stores s ON ft.user_id = s.user_id" +
		" INNER JOIN store_seats ss ON s.id = ss.store_id" +
		" WHERE ss.id = ?" +
		" ORDER BY ft.created_at DESC;"
	queryParams := []interface{}{seatId}

	if err := db.Raw(sqlQuery, queryParams...).Scan(&fcmTokens).Error; err != nil {
		return deviceTokens, err
	}

	for _, f := range fcmTokens {
		deviceTokens = append(deviceTokens, f.DeviceToken)
	}

	return deviceTokens, nil
}
