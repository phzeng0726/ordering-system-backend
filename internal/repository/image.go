package repository

import (
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

// 目前沒使用到
type ImagesRepo struct {
	db *gorm.DB
}

func NewImagesRepo(db *gorm.DB) *ImagesRepo {
	return &ImagesRepo{
		db: db,
	}
}

func (r *ImagesRepo) Create(tx *gorm.DB, imageBytes []byte) error {
	var data domain.Image
	data.BytesData = imageBytes

	if err := tx.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *ImagesRepo) GetById(tx *gorm.DB, imageId int) (domain.Image, error) {
	var data domain.Image

	if err := tx.Where("id = ?", imageId).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
