package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"

	"github.com/google/uuid"
)

type StoresService struct {
	repo repository.Stores
}

func NewStoresService(repo repository.Stores) *StoresService {
	return &StoresService{repo: repo}
}

func (s *StoresService) Create(ctx context.Context, store domain.Store) (string, error) {
	store.Id = uuid.New().String()
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

func (s *StoresService) GetByStoreId(ctx context.Context, userId string, storeId string) (domain.Store, error) {
	store, err := s.repo.GetByStoreId(ctx, userId, storeId)
	if err != nil {
		return store, err
	}
	return store, nil
}

func (s *StoresService) GetAll(ctx context.Context) ([]domain.Store, error) {
	stores, err := s.repo.GetAll(ctx)
	if err != nil {
		return stores, err
	}
	return stores, nil
}
