package service

import (
	"context"
	"fmt"

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

func (s itemServiceImpl) GetAllItems(ctx context.Context, limit int, cursor []interface{}) ([]domain.Item, []interface{}, error) {
	return s.repo.GetAllItems(ctx, limit, cursor)
}

func (s itemServiceImpl) AddItem(ctx context.Context, item domain.Item) (*domain.Item, error) {
	if item.Title == "" {
		return nil, fmt.Errorf("title cannot be nil")
	}

	if item.Code == "" {
		return nil, fmt.Errorf("code cannot be nil")
	}

	if item.Price <= 0 {
		return nil, fmt.Errorf("price cannot be zero")
	}

	if item.Stock > 0 {
		item.Available = true
	}

	itemNew, err := s.repo.AddItem(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	return itemNew, nil
}

func (s itemServiceImpl) ReadItem(ctx context.Context, id string) *domain.Item {
	item, _ := s.repo.ReadItem(ctx, id)

	return item
}

func (s itemServiceImpl) UpdateItem(ctx context.Context, itemNew domain.Item) *domain.Item {
	item, _ := s.repo.Update(ctx, itemNew)

	return item
}
