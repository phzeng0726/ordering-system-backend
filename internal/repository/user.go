package repository

import (
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

func (r *UsersRepo) Create(userId string, uq domain.UserRequest) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		client, err := firebase_auth.Init()
		if err != nil {
			return err
		}

		uidCode, err := firebase_auth.CreateUser(uq, client)
		if err != nil {
			if strings.Contains(err.Error(), "EMAIL_EXISTS") {
				return errors.New("email has already existed")
			}
			return err
		}

		userAccount := domain.UserAccount{
			Id:       userId,
			UidCode:  uidCode,
			Email:    uq.Email,
			UserType: uq.UserType,
		}

		user := domain.User{
			Id:         userId,
			FirstName:  uq.FirstName,
			LastName:   uq.LastName,
			LanguageId: uq.LanguageId,
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
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetByEmail(email string, userType int) (string, error) {
	var userAccount domain.UserAccount
	if err := r.db.Where("email = ?", email).Where("user_type = ?", userType).First(&userAccount).Error; err != nil {
		// 查無使用者，前端要收到false的消息
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userAccount.Id, nil
		}
		return userAccount.Id, err
	}

	fmt.Println(userAccount)
	return userAccount.Id, nil
}

func (r *UsersRepo) GetById(userId string) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", userId).Preload("UserAccount").Find(&user).Error; err != nil {
		return user, err
	}

	user.Email = user.UserAccount.Email
	return user, nil
}
