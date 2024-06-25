package repository

import (
	"time"

	"github.com/osalomon89/go-basics/internal/core/domain"
	"github.com/osalomon89/go-basics/internal/core/ports"
)

var items = []domain.Item{
	{
		ID:          1,
		Code:        "Item001",
		Title:       "Camisa",
		Description: "camisa de algodão",
		Price:       79,
		Stock:       3,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Code:        "Item002",
		Title:       "Bola futebol",
		Description: "bola futebol",
		Price:       20,
		Stock:       30,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

type itemRepositoryImpl struct {
}

func NewRepository() ports.ItemRepository {
	return itemRepositoryImpl{}
}

func (s itemRepositoryImpl) GetAllItems() []domain.Item {
	return items
}

func (s itemRepositoryImpl) AddItem(item domain.Item) []domain.Item {
	if item.Title == "" {
		return nil
	}

	items = append(items, item)

	return items
}

func (s itemRepositoryImpl) ReadItem(id int) *domain.Item {
	for _, item := range items {
		if id == item.ID {
			return &item
		}
	}

	return nil
}

func (s itemRepositoryImpl) UpdateItem(id int, itemNew domain.Item) *domain.Item {
	for i, v := range items {
		if id == v.ID {
			itemNew.ID = v.ID
			itemNew.CreatedAt = v.CreatedAt
			itemNew.UpdatedAt = time.Now()
			items[i] = itemNew

			return &itemNew
		}
	}

	return nil
}
