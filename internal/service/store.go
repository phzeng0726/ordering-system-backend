package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"

	"github.com/google/uuid"
)

type StoresService struct {
	repo         repository.Stores
	usersService UsersService
}

func NewStoresService(repo repository.Stores, usersService UsersService) *StoresService {
	return &StoresService{
		repo:         repo,
		usersService: usersService,
	}
}

func (s *StoresService) Create(ctx context.Context, store domain.Store) (string, error) {
	store.Id = uuid.New().String()

	// 確認該User存在，才可新增Store
	if _, err := s.usersService.GetById(ctx, store.UserId); err != nil {
		return store.Id, err
	}

	if err := s.repo.Create(ctx, store); err != nil {
		return store.Id, err
	}

	return store.Id, nil
}

func (s *StoresService) Update(ctx context.Context, store domain.Store) error {
	if err := s.repo.Update(ctx, store); err != nil {
		return err
	}

	return nil
}

func (s *StoresService) Delete(ctx context.Context, userId string, storeId string) error {
	if err := s.repo.Delete(ctx, userId, storeId); err != nil {
		return err
	}
	return nil
}

func (s *StoresService) GetAllByUserId(ctx context.Context, userId string) ([]domain.Store, error) {
	stores, err := s.repo.GetAllByUserId(ctx, userId)
	if err != nil {
		return stores, err
	}
	return stores, nil
}

func (s *StoresService) GetByStoreId(ctx context.Context, storeId string) (domain.Store, error) {
	store, err := s.repo.GetByStoreId(ctx, storeId)
	if err != nil {
		return store, err
	}
	return store, nil
}
