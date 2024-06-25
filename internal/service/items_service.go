package service

import (
	"time"

	"github.com/osalomon89/go-basics/internal/core/domain"
	"github.com/osalomon89/go-basics/internal/core/ports"
)

type itemServiceImpl struct {
	repo ports.ItemRepository
}

func NewService(repo ports.ItemRepository) ports.ItemService {
	return itemServiceImpl{
		repo: repo,
	}
}

func (s itemServiceImpl) GetAllItems() []domain.Item {
	return s.repo.GetAllItems()
}

func (s itemServiceImpl) AddItem(item domain.Item) []domain.Item {
	if item.Title == "" {
		return nil
	}

	return s.repo.AddItem(item)
}

func (s itemServiceImpl) ReadItem(id int) *domain.Item {
	items := s.repo.GetAllItems()
	for _, item := range items {
		if id == item.ID {
			return &item
		}
	}

	return nil
}

func (s itemServiceImpl) UpdateItem(id int, itemNew domain.Item) *domain.Item {
	items := s.repo.GetAllItems()

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
