package repository

import (
	"context"
	"io"
	"ordering-system-backend/internal/domain"
	"os"

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

func imageTesting() ([]byte, error) {
	var imageBytes []byte
	// 開啟圖片
	file, err := os.Open("C:/Users/phzen/Desktop/LINE_ALBUM_Lucky_230917_1.jpg")
	if err != nil {

		return imageBytes, err
	}
	defer file.Close()

	// 轉為 bytes
	imageBytes, err = io.ReadAll(file)
	if err != nil {
		return imageBytes, err
	}

	return imageBytes, nil
}

func (r *SeatsRepo) GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error) {
	var seat domain.Seat
	var data domain.BinaryData
	bytes, err := imageTesting()
	if err != nil {
		return seat, err
	}

	data.BinData = bytes
	db := r.db.WithContext(ctx)

	if err := db.Create(&data).Error; err != nil {
		return seat, err
	}

	return seat, nil
}

// func (r *SeatsRepo) GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error) {
// 	var seat domain.Seat
// 	db := r.db.WithContext(ctx)

// 	if err := db.Where("store_id = ? AND id = ?", storeId, seatId).First(&seat).Error; err != nil {
// 		return seat, err
// 	}

// 	return seat, nil
// }
