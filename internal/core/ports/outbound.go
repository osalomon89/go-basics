package ports

import (
	"context"

	"github.com/osalomon89/go-basics/internal/core/domain"
)

type ItemRepository interface {
	GetAllItems() []domain.Item
	AddItem(ctx context.Context, item domain.Item) (*domain.Item, error)
	ReadItem(ctx context.Context, id string) (*domain.Item, error)
	Update(ctx context.Context, itemNew domain.Item) (*domain.Item, error)
	Delete(ctx context.Context, id string) error
}
