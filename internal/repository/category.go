package repository

import (
	"context"
	"errors"
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

func (r *CategoriesRepo) Update(ctx context.Context, userId string, category domain.Category, categoryLanguage domain.CategoryLanguage) error {
	var tempCategory domain.Category
	var tempCategoryUserMappings []domain.CategoryUserMapping
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", category.Id).First(&tempCategory).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("category not found")
			}
			return err
		}

		// 系統預設category無法update
		if *tempCategory.IsDefault {
			return errors.New("cannot update default category")
		}

		// 確認該category屬於該user
		if err := tx.Where("category_id = ? AND user_id = ?", category.Id, userId).First(&tempCategoryUserMappings).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("category id not exist with user id")
			}
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

func (r *CategoriesRepo) Delete(ctx context.Context, categoryId int) error {
	return nil
}

func (r *CategoriesRepo) GetAllByUserId(ctx context.Context, userId string, languageId int) ([]domain.Category, error) {
	type result struct {
		Id         int
		Identifier string
		IsDefault  *bool
		Title      string
	}
	var res []result
	var categories []domain.Category
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := r.rt.CheckUserExist(tx, userId); err != nil {
			return err
		}

		sqlQuery := "SELECT c.*, cl.title" +
			" FROM category_user_mapping cum" +
			" INNER JOIN categories c on cum.category_id = c.id" +
			" INNER JOIN category_language cl on c.id = cl.category_id" +
			" WHERE cum.user_id = ?" +
			" AND cl.language_id = ? OR cl.language_id IS NULL;"
		queryParams := []interface{}{userId, languageId}

		if err := tx.Raw(sqlQuery, queryParams...).Scan(&res).Error; err != nil {
			return err
		}

		for _, r := range res {
			categories = append(categories, domain.Category{
				Id:         r.Id,
				Identifier: r.Identifier,
				IsDefault:  r.IsDefault,
				Title:      r.Title,
			})
		}

		return nil
	}); err != nil {
		return categories, err
	}

	return categories, nil
}
