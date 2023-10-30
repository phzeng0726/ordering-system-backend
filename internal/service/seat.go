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

func (s *SeatsService) GetById(ctx context.Context, storeId string, seatId int) (domain.Seat, error) {
	seat, err := s.repo.GetById(ctx, storeId, seatId)
	if err != nil {
		return seat, err
	}

	return seat, nil
}
