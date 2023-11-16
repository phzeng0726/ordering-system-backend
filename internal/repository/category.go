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

func (r *CategoriesRepo) Create(ctx context.Context, category domain.Category, categoryLanguage domain.CategoryLanguage, categoryUserMapping domain.CategoryUserMapping) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&category).Error; err != nil {
			return err
		}

		categoryLanguage.CategoryId = category.Id
		categoryUserMapping.CategoryId = category.Id

		if err := tx.Create(&categoryLanguage).Error; err != nil {
			return err
		}

		if err := tx.Create(&categoryUserMapping).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepo) Update(ctx context.Context, category domain.Category, categoryLanguage domain.CategoryLanguage, categoryUserMapping domain.CategoryUserMapping) error {
	// db := r.db.WithContext(ctx)

	// if err := db.Transaction(func(tx *gorm.DB) error {
	// 	if err := tx.Model(&domain.Menu{}).Where("id = ?", menu.Id).Updates(updatedMenu).Error; err != nil {
	// 		return err
	// 	}
	// 	// if err := tx.Create(&category).Error; err != nil {
	// 	// 	return err
	// 	// }

	// 	// categoryLanguage.CategoryId = category.Id
	// 	// categoryUserMapping.CategoryId = category.Id

	// 	// if err := tx.Create(&categoryLanguage).Error; err != nil {
	// 	// 	return err
	// 	// }

	// 	// if err := tx.Create(&categoryUserMapping).Error; err != nil {
	// 	// 	return err
	// 	// }

	// 	return nil
	// }); err != nil {
	// 	return err
	// }

	return nil
}

func (r *CategoriesRepo) Delete(ctx context.Context, categoryId int) error {
	return nil
}

func (r *CategoriesRepo) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.CategoryUserMapping, error) {
	var categoryUserMappings []domain.CategoryUserMapping
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckUserExist(tx, userId); err != nil {
			return err
		}

		if err := tx.Preload("Category").
			Preload("Category.CategoryLanguage", "language_id = ? OR language_id IS NULL", languageId).
			Where("user_id = ?", userId).
			Find(&categoryUserMappings).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return categoryUserMappings, err
	}

	return categoryUserMappings, nil
}
