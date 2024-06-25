package ports

import "github.com/osalomon89/go-basics/internal/core/domain"

type ItemService interface {
	GetAllItems() []domain.Item
	AddItem(item domain.Item) []domain.Item
	ReadItem(id int) *domain.Item
	UpdateItem(id int, itemNew domain.Item) *domain.Item
}
