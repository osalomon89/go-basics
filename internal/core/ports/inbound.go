package ports

import (
	"context"

	"github.com/osalomon89/go-basics/internal/core/domain"
)

type ItemService interface {
	GetAllItems() []domain.Item
	AddItem(ctx context.Context, item domain.Item) (*domain.Item, error)
	ReadItem(ctx context.Context, id string) *domain.Item
	UpdateItem(ctx context.Context, itemNew domain.Item) *domain.Item
}
