package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
	"strings"

	"github.com/google/uuid"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(ctx context.Context, input CreateUserInput) (string, error) {
	userId := uuid.New().String()  // DB
	uidCode := uuid.New().String() // Firebase Auth
	upperUidCode := strings.ToUpper(uidCode)
	lowerEmail := strings.ToLower(input.Email)

	userAccount := domain.UserAccount{
		Id:       userId,
		UidCode:  upperUidCode,
		Email:    lowerEmail,
		UserType: input.UserType,
	}

	user := domain.User{
		Id:         userId,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		LanguageId: input.LanguageId,
	}

	if err := s.repo.Create(ctx, userAccount, user, input.Password); err != nil {
		return userId, err
	}
	return userId, nil
}

func (s *UsersService) Update(ctx context.Context, userId string, input UpdateUserInput) error {
	newUser := domain.User{
		Id:         userId,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		LanguageId: input.LanguageId,
	}

	if err := s.repo.Update(ctx, newUser); err != nil {
		return err
	}
	return nil
}

func (s *UsersService) Delete(ctx context.Context, userId string) error {
	if err := s.repo.Delete(ctx, userId); err != nil {
		return err
	}
	return nil
}

func (s *UsersService) GetByEmail(ctx context.Context, email string, userType int) (string, error) {
	lowerEmail := strings.ToLower(email)

	userId, err := s.repo.GetByEmail(ctx, lowerEmail, userType)
	if err != nil {
		return userId, err
	}
	return userId, nil
}

func (s *UsersService) GetByUid(ctx context.Context, uid string, userType int) (string, error) {
	userId, err := s.repo.GetByUid(ctx, uid, userType)
	if err != nil {
		return userId, err
	}
	return userId, nil
}

func (s *UsersService) GetById(ctx context.Context, userId string) (domain.User, error) {
	user, err := s.repo.GetById(ctx, userId)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UsersService) ResetPassword(ctx context.Context, input ResetPasswordInput) error {
	if err := s.repo.ResetPassword(ctx, input.UserId, input.Password); err != nil {
		return err
	}
	return nil
}
