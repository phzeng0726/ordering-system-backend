package repository

import (
	"context"
	"errors"
	"fmt"
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

		err = firebase_auth.CreateUser(client, userAccount.Email, password, userAccount.UidCode)
		if err != nil {
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

func (r *UsersRepo) Update(ctx context.Context, u domain.User) error {
	var user domain.User
	res := r.db.WithContext(ctx).Model(&user).Where("id = ?", u.Id).Updates(&u)
	if res.Error != nil {
		return res.Error
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

	fmt.Println(userAccount)
	return userAccount.Id, nil
}

func (r *UsersRepo) GetById(ctx context.Context, userId string) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("id = ?", userId).Preload("UserAccount").Find(&user).Error; err != nil {
		return user, err
	}

	user.Email = user.UserAccount.Email
	return user, nil
}

func (r *UsersRepo) Delete(ctx context.Context, userId string) error {
	var user domain.User
	var userAccount domain.UserAccount

	res := r.db.WithContext(ctx).Where("id = ?", userId).First(&userAccount)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user id not found")
		}
		return res.Error
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		client, err := firebase_auth.Init()
		if err != nil {
			return err
		}

		if err := tx.Where("id = ?", userId).Delete(&userAccount).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", userId).Delete(&user).Error; err != nil {
			return err
		}

		if err := firebase_auth.DeleteUser(userAccount.UidCode, client); err != nil {
			return errors.New("firebase uid code not found in firebase")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) ResetPassword(ctx context.Context, ur domain.UserRequest) error {
	var userAccount domain.UserAccount
	res := r.db.WithContext(ctx).Where("id = ?", ur.UserId).First(&userAccount)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user id not found")
		}
		return res.Error
	}

	client, err := firebase_auth.Init()
	if err != nil {
		return err
	}

	if err = firebase_auth.ResetPassword(ur, userAccount.UidCode, client); err != nil {
		return err
	}

	return nil
}
