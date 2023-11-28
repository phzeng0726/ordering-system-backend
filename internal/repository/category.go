package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type CategoriesRepo struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
	}
}

func (r *CategoriesRepo) checkCategoryPermission(tx *gorm.DB, userId string, categoryId int) error {
	var tempCategory domain.Category
	var tempCategoryUserMapping domain.CategoryUserMapping

	// 確認category存在
	if err := tx.Where("id = ?", categoryId).First(&tempCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return err
	}

	// 系統預設的category無法變更、刪除
	if *tempCategory.IsDefault {
		return errors.New("default category cannot be edit")
	}

	// 確認該category屬於該user
	if err := tx.Where("category_id = ? AND user_id = ?", categoryId, userId).First(&tempCategoryUserMapping).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category id not exist with user id")
		}
		return err
	}

	return nil
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

func (r *CategoriesRepo) Update(ctx context.Context, userId string, category domain.Category, categoryLanguage domain.CategoryLanguage) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 確認category存在、非default，且為該user所擁有
		if err := r.checkCategoryPermission(tx, userId, category.Id); err != nil {
			return err
		}

		if err := tx.Model(&domain.Category{}).Where("id = ?", category.Id).Updates(&category).Error; err != nil {
			return err
		}

		if err := tx.Model(&domain.CategoryLanguage{}).Where("category_id = ?", category.Id).Updates(&categoryLanguage).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepo) Delete(ctx context.Context, userId string, categoryId int) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 確認category存在、非default，且為該user所擁有
		if err := r.checkCategoryPermission(tx, userId, categoryId); err != nil {
			return err
		}

		if err := tx.Where("category_id = ?", categoryId).Delete(&domain.CategoryUserMapping{}).Error; err != nil {

			return err
		}

		if err := tx.Where("category_id = ?", categoryId).Delete(&domain.CategoryLanguage{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", categoryId).Delete(&domain.Category{}).Error; err != nil {
			if strings.Contains(err.Error(), "menu_items_ibfk_1") {
				return errors.New("category already in use by menu items")
			}
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *CategoriesRepo) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.CategoryUserMapping, error) {
	var categoryUserMappings []domain.CategoryUserMapping
	db := r.db.WithContext(ctx)

	if err := db.Preload("Category").
		Preload("Category.CategoryLanguage", "language_id IS NULL OR language_id = ?", languageId).
		Where("user_id = ?", userId).
		Find(&categoryUserMappings).Error; err != nil {
		return categoryUserMappings, err
	}

	return categoryUserMappings, nil
}

// GetAllByUserId 也可寫這樣
// SELECT c.*, cl.title
// FROM category_user_mapping cum
// INNER JOIN categories c on cum.category_id = c.id
// INNER JOIN category_language cl on c.id = cl.category_id
// WHERE cum.user_id = ?
// AND (cl.language_id IS NULL OR cl.language_id = ?);
