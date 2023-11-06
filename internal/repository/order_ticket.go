package repository

import (
	"context"
	"errors"
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

func (r *OrderTicketsRepo) Create(ctx context.Context, ticket domain.OrderTicket) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 因為有設fKey，會自動幫忙create orderItems
		if err := tx.Create(&ticket).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *OrderTicketsRepo) Update(ctx context.Context, storeId string, ticket domain.OrderTicket) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.OrderTicket{}).Where("id = ?", ticket.Id).Updates(ticket).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *OrderTicketsRepo) Delete(ctx context.Context, storeId string, ticketId int) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", ticketId).First(&domain.OrderTicket{}).Error; err != nil {
			return errors.New("order ticket not found")
		}

		if err := tx.Where("order_id = ?", ticketId).Delete(&domain.OrderTicketItem{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", ticketId).Delete(&domain.OrderTicket{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *OrderTicketsRepo) GetAllByStoreId(ctx context.Context, storeId string) ([]domain.OrderTicket, error) {
	var orderTickets []domain.OrderTicket
	var seatIds []int
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 獲取該商店所有的seatId
		querySeatIds := "SELECT ss.id" +
			" FROM stores s" +
			" JOIN store_seats ss ON ss.store_id = s.id" +
			" WHERE s.id = ?"

		if err := tx.Raw(querySeatIds, storeId).Scan(&seatIds).Error; err != nil {
			return err
		}

		if err := tx.Preload("OrderTicketItems").
			Where("seat_id IN (?)", seatIds).
			Find(&orderTickets).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return orderTickets, err
	}

	return orderTickets, nil
}
