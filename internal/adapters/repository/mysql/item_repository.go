package mysqlrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/osalomon89/go-basics/internal/core/domain"
	"github.com/osalomon89/go-basics/internal/core/ports"
)

var items = []domain.Item{
	{
		ID:          "1",
		Code:        "Item001",
		Title:       "Camisa",
		Description: "camisa de algodão",
		Price:       79,
		Stock:       3,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	},
	{
		ID:          "2",
		Code:        "Item002",
		Title:       "Bola futebol",
		Description: "bola futebol",
		Price:       20,
		Stock:       30,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	},
}

type itemRepositoryImpl struct {
}

func NewMySQLRepository() ports.ItemRepository {
	return itemRepositoryImpl{}
}

func (s itemRepositoryImpl) GetAllItems(ctx context.Context, limit int, searchAfter []interface{}) ([]domain.Item, []interface{}, error) {
	return items, nil, nil
}

func (s itemRepositoryImpl) AddItem(ctx context.Context, item domain.Item) (*domain.Item, error) {
	for _, i := range items {
		if i.Code == item.Code {
			return nil, fmt.Errorf("duplicated entry")
		}
	}

	item.ID = string((len(items) + 1))
	now := time.Now()
	item.CreatedAt = &now
	item.UpdatedAt = &now

	items = append(items, item)

	return &item, nil
}

func (s itemRepositoryImpl) ReadItem(ctx context.Context, id string) (*domain.Item, error) {
	for _, item := range items {
		if id == item.ID {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("item not found")
}

func (s itemRepositoryImpl) Update(ctx context.Context, itemNew domain.Item) (*domain.Item, error) {
	for i, v := range items {
		if itemNew.ID == v.ID {
			itemNew.ID = v.ID
			itemNew.CreatedAt = v.CreatedAt
			now := time.Now()
			itemNew.UpdatedAt = &now
			items[i] = itemNew

			return &itemNew, nil
		}
	}

	return nil, fmt.Errorf("item not found")
}

func (s itemRepositoryImpl) Delete(ctx context.Context, id string) error {
	for i, v := range items {
		if id == v.ID {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("item not found")
}
