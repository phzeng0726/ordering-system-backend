package repository

import (
	"context"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type OrderTicketsRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewOrderTicketsRepo(db *gorm.DB, rt *RepoTools) *OrderTicketsRepo {
	return &OrderTicketsRepo{
		db: db,
		rt: rt,
	}
}

func (r *OrderTicketsRepo) createOrderItems(tx *gorm.DB, orderId int, orderItems []domain.OrderTicketItem) error {
	for _, oi := range orderItems {
		oi.OrderId = orderId

		if err := tx.Create(&oi).Error; err != nil {
			return err
		}

	}

	return nil
}

func (r *OrderTicketsRepo) Create(ctx context.Context, ticket domain.OrderTicket) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&ticket).Error; err != nil {
			return err
		}

		if err := r.createOrderItems(tx, ticket.Id, ticket.OrderTicketItems); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}
