package repository

import (
	"context"
	"errors"
	"log"
	"ordering-system-backend/internal/domain"
	firebase_auth "ordering-system-backend/pkg/firebase_core/auth"
	"strings"

	"firebase.google.com/go/auth"
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

// 只有這個file會用到，所以用private就夠了
func (r *UsersRepo) getUserAccountById(tx *gorm.DB, userId string) (domain.UserAccount, error) {
	var userAccount domain.UserAccount

	if err := tx.Where("id = ?", userId).First(&userAccount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userAccount, errors.New("user id not found")
		}
		return userAccount, err
	}

	return userAccount, nil
}

func (r *UsersRepo) createUser(tx *gorm.DB, userAccount *domain.UserAccount, user *domain.User) error {
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
}

func (r *UsersRepo) createCategories(tx *gorm.DB, userId string) error {
	var categories []domain.Category

	if err := tx.Where("is_default = ?", true).Find(&categories).Error; err != nil {
		return err
	}

	for _, c := range categories {
		var categoryUserMapping domain.CategoryUserMapping
		categoryUserMapping.UserId = userId
		categoryUserMapping.CategoryId = c.Id

		if err := tx.Create(&categoryUserMapping).Error; err != nil {
			return errors.New("failed to create default categories: " + err.Error())
		}
	}

	return nil
}

func (r *UsersRepo) Create(ctx context.Context, userAccount domain.UserAccount, user domain.User, password string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		client, err := firebase_auth.Init(*userAccount.UserType)
		if err != nil {
			return err
		}

		if err := r.createUser(tx, &userAccount, &user); err != nil {
			return err
		}

		if err := r.createCategories(tx, user.Id); err != nil {
			return err
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
		if _, err := r.getUserAccountById(tx, user.Id); err != nil {
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

func (r *UsersRepo) deleteMenus(tx *gorm.DB, userId string) error {
	type result struct {
		MenuId     string
		MenuItemId int
	}
	var res []result
	var menuIds []string
	var menuItemIds []int

	sqlQuery := "SELECT mim.menu_id, mim.menu_item_id" +
		" FROM menus m" +
		" INNER JOIN menu_item_mapping mim ON mim.menu_id = m.id" +
		" INNER JOIN menu_items mi ON mi.id = mim.menu_item_id" +
		" WHERE m.user_id = ?;"
	queryParams := []interface{}{userId}

	if err := tx.Raw(sqlQuery, queryParams...).Scan(&res).Error; err != nil {
		return err
	}

	for _, item := range res {
		menuIds = append(menuIds, item.MenuId)
		menuItemIds = append(menuItemIds, item.MenuItemId)
	}

	if err := tx.Where("menu_id IN (?)", menuIds).Delete(&domain.MenuItemMapping{}).Error; err != nil {
		return err
	}

	if err := tx.Where("id IN (?)", menuItemIds).Delete(&domain.MenuItem{}).Error; err != nil {
		return err
	}

	// 查詢和刪除相關的 store_menu_mapping
	if err := tx.Where("store_id IN (SELECT id FROM stores WHERE user_id = ?)", userId).Delete(&domain.StoreMenuMapping{}).Error; err != nil {
		return err
	}

	if err := tx.Where("user_id = ?", userId).Delete(&domain.Menu{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) deleteStores(tx *gorm.DB, userId string) error {
	// 查詢和刪除相關的 store_opening_hours
	if err := tx.Where("store_id IN (SELECT id FROM stores WHERE user_id = ?)", userId).Delete(&domain.StoreOpeningHour{}).Error; err != nil {
		return err
	}

	// 查詢和刪除相關的 store_seats
	if err := tx.Where("store_id IN (SELECT id FROM stores WHERE user_id = ?)", userId).Delete(&domain.Seat{}).Error; err != nil {
		return err
	}

	// 刪除 stores 表中具有特定 user_id 的行
	if err := tx.Where("user_id = ?", userId).Delete(&domain.Store{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) deleteCategories(tx *gorm.DB, userId string) error {
	if err := tx.Where("user_id = ?", userId).Delete(&domain.CategoryUserMapping{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) deleteFCMTokens(tx *gorm.DB, userId string) error {
	if err := tx.Where("user_id = ?", userId).Delete(&domain.FCMToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) deleteUser(tx *gorm.DB, client *auth.Client, userAccount domain.UserAccount, userId string) error {
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
}

func (r *UsersRepo) Delete(ctx context.Context, userId string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		userAccount, err := r.getUserAccountById(tx, userId)
		if err != nil {
			return err
		}

		client, err := firebase_auth.Init(*userAccount.UserType)
		if err != nil {
			return err
		}

		log.Println("Delete menu")
		if err := r.deleteMenus(tx, userId); err != nil {
			return err
		}

		log.Println("Delete store")
		if err := r.deleteStores(tx, userId); err != nil {
			return err
		}

		log.Println("Delete category")
		if err := r.deleteCategories(tx, userId); err != nil {
			return err
		}

		log.Println("Delete fcm token")
		if err := r.deleteFCMTokens(tx, userId); err != nil {
			return err
		}

		log.Println("Delete user")
		if err := r.deleteUser(tx, client, userAccount, userId); err != nil {
			return err
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

	if err := db.Where("email = ? AND user_type = ?", email, userType).First(&userAccount).Error; err != nil {
		// 查無使用者，前端要收到false的消息
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userAccount.Id, nil
		}
		return userAccount.Id, err
	}
	return userAccount.Id, nil
}

// By firebase uid
func (r *UsersRepo) GetByUid(ctx context.Context, uid string, userType int) (string, error) {
	var userAccount domain.UserAccount
	db := r.db.WithContext(ctx)

	if err := db.Where("uid_code = ? AND user_type = ?", uid, userType).First(&userAccount).Error; err != nil {
		// 查無使用者，前端要收到false的消息
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userAccount.Id, nil
		}
		return userAccount.Id, err
	}
	return userAccount.Id, nil
}

// User
func (r *UsersRepo) GetById(ctx context.Context, userId string) (domain.User, error) {
	var user domain.User
	db := r.db.WithContext(ctx)

	if err := db.Where("id = ?", userId).Preload("UserAccount").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user id not found")
		}
		return user, err
	}

	user.Email = user.UserAccount.Email
	return user, nil
}

func (r *UsersRepo) ResetPassword(ctx context.Context, userId string, newPassword string) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		userAccount, err := r.getUserAccountById(tx, userId)
		if err != nil {
			return err
		}

		client, err := firebase_auth.Init(*userAccount.UserType)
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
