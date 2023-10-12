package repository

import (
	"context"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type CategoriesRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewCategoriesRepo(db *gorm.DB, rt *RepoTools) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
		rt: rt,
	}
}

func (r *CategoriesRepo) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Category, error) {
	var categories []domain.Category
	var categoryUserMappings []domain.CategoryUserMapping

	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckUserExist(tx, userId); err != nil {
			return err
		}

		if err := tx.Preload("Category").
			Preload("Category.CategoryLanguage", "language_id = ?", languageId).
			Where("user_id = ?", userId).
			Find(&categoryUserMappings).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return categories, err
	}

	for _, cum := range categoryUserMappings {
		categories = append(categories, cum.Category)
	}

	return categories, nil
}
