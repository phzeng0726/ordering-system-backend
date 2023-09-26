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
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) getUserAccountFromDB(ctx context.Context, userId string) (domain.UserAccount, error) {
	var userAccount domain.UserAccount

	res := r.db.WithContext(ctx).Where("id = ?", userId).First(&userAccount)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return userAccount, errors.New("user id not found")
		}
		return userAccount, res.Error
	}

	return userAccount, res.Error
}

func (r *UsersRepo) Create(ctx context.Context, userAccount domain.UserAccount, user domain.User, password string) error {
	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
	if _, err := r.getUserAccountFromDB(ctx, user.Id); err != nil {
		return err
	}

	res := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", user.Id).Updates(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *UsersRepo) Delete(ctx context.Context, userId string) error {
	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		userAccount, err := r.getUserAccountFromDB(ctx, userId)
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
	if err := r.db.WithContext(ctx).Where("email = ?", email).Where("user_type = ?", userType).First(&userAccount).Error; err != nil {
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
	if err := r.db.WithContext(ctx).Where("id = ?", userId).Preload("UserAccount").First(&user).Error; err != nil {
		return user, err
	}

	user.Email = user.UserAccount.Email
	return user, nil
}

func (r *UsersRepo) ResetPassword(ctx context.Context, userId string, newPassword string) error {
	var userAccount domain.UserAccount
	userAccount, err := r.getUserAccountFromDB(ctx, userId)
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
}
