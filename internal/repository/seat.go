package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type SeatsRepo struct {
	db *gorm.DB
}

func NewSeatsRepo(db *gorm.DB) *SeatsRepo {
	return &SeatsRepo{
		db: db,
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

func (r *SeatsRepo) checkSeatExist(tx *gorm.DB, seatId int) error {
	if err := tx.Where("id = ?", seatId).First(&domain.Seat{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("seat id not found")
		}
		return err
	}

	return nil
}

func (r *SeatsRepo) Delete(ctx context.Context, storeId string, seatId int) error {
	db := r.db.WithContext(ctx)
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.checkSeatExist(tx, seatId); err != nil {
			return err
		}

		if err := tx.Where("id = ?", seatId).Delete(&domain.Seat{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *SeatsRepo) GetAllByStoreId(ctx context.Context, storeId string) ([]domain.Seat, error) {
	var seats []domain.Seat
	db := r.db.WithContext(ctx)

	if err := db.Where("store_id = ?", storeId).Find(&seats).Error; err != nil {
		return seats, err
	}

	return seats, nil
}

func (r *SeatsRepo) GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error) {
	var seat domain.Seat
	db := r.db.WithContext(ctx)

	if err := db.Where("store_id = ? AND id = ?", storeId, seatId).First(&seat).Error; err != nil {
		return seat, err
	}

	return seat, nil
}

func (r *SeatsRepo) GetSeatBySeatId(ctx context.Context, seatId int) (domain.Seat, error) {
	var seat domain.Seat
	db := r.db.WithContext(ctx)

	if err := db.Where("id = ?", seatId).First(&seat).Error; err != nil {
		return seat, err
	}

	return seat, nil
}
