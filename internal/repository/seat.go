package repository

import (
	"context"
	"fmt"
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

func loadImage() ([]byte, error) {
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

func (r *SeatsRepo) insertImage(ctx context.Context) error {
	var data domain.BinaryData
	bytes, err := loadImage()
	if err != nil {
		return err
	}

	data.BinData = bytes
	db := r.db.WithContext(ctx)

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *SeatsRepo) LoadImage() (domain.BinaryData, error) {
	var data domain.BinaryData

	if err := r.db.Where("id = ?", 1).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *SeatsRepo) GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error) {
	var seat domain.Seat
	// if err := r.insertImage(ctx); err != nil {
	// 	return seat, err
	// }

	data, err := r.LoadImage()
	if err != nil {
		return seat, err
	}

	fmt.Println(data)
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
