package repository

import (
	"errors"
	"fmt"
	"ordering-system-backend/domain"

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

func (r *UsersRepo) Create(userId string, u domain.UserRequest) error {
	// 如果Firebase uid已經存在於DB，報錯
	var userAccounts []domain.UserAccount
	if err := r.db.Where("uid_code = ?", u.UidCode).Find(&userAccounts).Error; err != nil {
		return err
	}

	if len(userAccounts) != 0 {
		return errors.New("firebase uid already exist")
	}

	userAccount := domain.UserAccount{
		Id:       userId,
		UidCode:  u.UidCode,
		Email:    u.Email,
		UserType: u.UserType,
	}

	user := domain.User{
		Id:         userId,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		LanguageId: u.LanguageId,
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userAccount).Error; err != nil {
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
