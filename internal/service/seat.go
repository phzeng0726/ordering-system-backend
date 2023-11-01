package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

type SeatsService struct {
	repo repository.Seats
}

func NewSeatsService(repo repository.Seats) *SeatsService {
	return &SeatsService{repo: repo}
}

func (s *SeatsService) Create(ctx context.Context, seat domain.Seat) error {
	if err := s.repo.Create(ctx, seat); err != nil {
		return err
	}

	return nil
}

func (s *SeatsService) Update(ctx context.Context, seat domain.Seat) error {
	if err := s.repo.Update(ctx, seat); err != nil {
		return err
	}

	return nil
}

func (s *SeatsService) Delete(ctx context.Context, storeId string, seatId int) error {
	if err := s.repo.Delete(ctx, storeId, seatId); err != nil {
		return err
	}

	return nil
}

func (s *SeatsService) GetAllByStoreId(ctx context.Context, storeId string) ([]domain.Seat, error) {
	seats, err := s.repo.GetAllByStoreId(ctx, storeId)
	if err != nil {
		return seats, err
	}

	return seats, nil
}

func (s *SeatsService) GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error) {
	seat, err := s.repo.GetById(ctx, storeId, seatId)
	if err != nil {
		return seat, err
	}

	return seat, nil
}
