package service

import (
	"fmt"
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

func (s itemServiceImpl) AddItem(item domain.Item) (*domain.Item, error) {
	if item.Title == "" {
		return nil, fmt.Errorf("title cannot be nil")
	}

	if item.Code == "" {
		return nil, fmt.Errorf("code cannot be nil")
	}

	if item.Price <= 0 {
		return nil, fmt.Errorf("price cannot be zero")
	}

	itemNew, err := s.repo.AddItem(item)
	if err != nil {
		return nil, fmt.Errorf("error in repository")
	}

	return itemNew, nil
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
