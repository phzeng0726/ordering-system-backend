package repository

import (
	"context"
	"errors"
	"fmt"
	"ordering-system-backend/domain"
	"strings"

	"gorm.io/gorm"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func firebaseInit() (*auth.Client, error) {
	opt := option.WithCredentialsFile("firebase_credential.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	// Access Auth service from default app
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func createFirebaseUser(uq domain.UserRequest, client *auth.Client) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(uq.Email).
		Password("secretPassword").
		DisplayName(uq.LastName + " " + uq.FirstName)

	u, err := client.CreateUser(context.Background(), params)

	if err != nil {
		return "", err
	}

	fmt.Printf("Successfully created user: %v\n", u)
	return u.UID, nil
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) Create(userId string, uq domain.UserRequest) error {
	// 如果Firebase uid已經存在於DB，報錯
	var userAccounts []domain.UserAccount
	if err := r.db.Where("uid_code = ?", uq.UidCode).Find(&userAccounts).Error; err != nil {
		return err
	}

	if len(userAccounts) != 0 {
		return errors.New("firebase uid already exist")
	}

	userAccount := domain.UserAccount{
		Id:       userId,
		UidCode:  uq.UidCode,
		Email:    uq.Email,
		UserType: uq.UserType,
	}

	user := domain.User{
		Id:         userId,
		FirstName:  uq.FirstName,
		LastName:   uq.LastName,
		LanguageId: uq.LanguageId,
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		client, err := firebaseInit()
		if err != nil {
			return err
		}

		if _, err := createFirebaseUser(uq, client); err != nil {
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
