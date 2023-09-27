package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"
	firebase_auth "ordering-system-backend/pkg/auth"
	"strings"

	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewUsersRepo(db *gorm.DB, rt *RepoTools) *UsersRepo {
	return &UsersRepo{
		db: db,
		rt: rt,
	}
}

func (r *UsersRepo) Create(ctx context.Context, userAccount domain.UserAccount, user domain.User, password string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		client, err := firebase_auth.Init()
		if err != nil {
			return err
		}

		if err := tx.Create(&userAccount).Error; err != nil {
			if strings.Contains(err.Error(), "unique_email_user_type") {
				return errors.New("email has already existed")
			}
			return errors.New("failed to create user account: " + err.Error())
		}

		if err := tx.Create(&user).Error; err != nil {
			return errors.New("failed to create user: " + err.Error())
		}

		if err := firebase_auth.CreateUser(client, userAccount.UidCode, userAccount.Email, password); err != nil {
			if strings.Contains(err.Error(), "EMAIL_EXISTS") {
				return errors.New("email has already existed in firebase")
			}
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) Update(ctx context.Context, user domain.User) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		_, err := r.rt.GetUserAccount(tx, user.Id)
		if err != nil {
			return err
		}

		if err := tx.Model(&domain.User{}).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) Delete(ctx context.Context, userId string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		userAccount, err := r.rt.GetUserAccount(tx, userId)
		if err != nil {
			return err
		}

		client, err := firebase_auth.Init()
		if err != nil {
			return err
		}

		// 查詢和刪除相關的 store_opening_hours
		if err := tx.Where("store_id IN (SELECT id FROM stores WHERE user_id = ?)", userId).Delete(&domain.StoreOpeningHour{}).Error; err != nil {
			return err
		}

		// 刪除 stores 表中具有特定 user_id 的行
		if err := tx.Where("user_id = ?", userId).Delete(&domain.Store{}).Error; err != nil {
			return err
		}

		// 刪除 user
		if err := tx.Where("id = ?", userId).Delete(&domain.User{}).Error; err != nil {
			return err
		}

		// 刪除 user_account
		if err := tx.Where("id = ?", userId).Delete(&userAccount).Error; err != nil {
			return err
		}

		if err := firebase_auth.DeleteUser(userAccount.UidCode, client); err != nil {
			return errors.New("firebase uid code not found in firebase")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetByEmail(ctx context.Context, email string, userType int) (string, error) {
	var userAccount domain.UserAccount
	db := r.db.WithContext(ctx)

	if err := db.Where("email = ?", email).Where("user_type = ?", userType).First(&userAccount).Error; err != nil {
		// 查無使用者，前端要收到false的消息
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userAccount.Id, nil
		}
		return userAccount.Id, err
	}
	return userAccount.Id, nil
}

func (r *UsersRepo) GetById(ctx context.Context, userId string) (domain.User, error) {
	var user domain.User
	db := r.db.WithContext(ctx)

	if err := db.Where("id = ?", userId).Preload("UserAccount").First(&user).Error; err != nil {
		return user, err
	}

	user.Email = user.UserAccount.Email
	return user, nil
}

func (r *UsersRepo) ResetPassword(ctx context.Context, userId string, newPassword string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		userAccount, err := r.rt.GetUserAccount(tx, userId)
		if err != nil {
			return err
		}

		client, err := firebase_auth.Init()
		if err != nil {
			return err
		}

		if err = firebase_auth.ResetPassword(client, userAccount.UidCode, newPassword); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
