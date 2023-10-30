package repository

import (
	"context"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type SeatsRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewSeatsRepo(db *gorm.DB, rt *RepoTools) *SeatsRepo {
	return &SeatsRepo{
		db: db,
		rt: rt,
	}
}

func (r *SeatsRepo) Create(ctx context.Context, seat domain.Seat) error {
	db := r.db.WithContext(ctx)

	if err := db.Create(&seat).Error; err != nil {
		return err
	}

	return nil
}

func (r *SeatsRepo) Update(ctx context.Context, seat domain.Seat) error {
	db := r.db.WithContext(ctx)

	if err := db.Model(&domain.Seat{}).Where("id = ?", seat.Id).Updates(&seat).Error; err != nil {
		return err
	}

	return nil
}

func (r *SeatsRepo) GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error) {
	var seat domain.Seat
	db := r.db.WithContext(ctx)

	if err := db.Where("store_id = ? AND id = ?", storeId, seatId).First(&seat).Error; err != nil {
		return seat, err
	}

	return seat, nil
}
