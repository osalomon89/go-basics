package mysqlrepo

// import (
// 	"fmt"
// 	"time"

// 	"github.com/osalomon89/go-basics/internal/core/domain"
// 	"github.com/osalomon89/go-basics/internal/core/ports"
// )

// var items = []domain.Item{
// 	{
// 		ID:          "1",
// 		Code:        "Item001",
// 		Title:       "Camisa",
// 		Description: "camisa de algodão",
// 		Price:       79,
// 		Stock:       3,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	},
// 	{
// 		ID:          2,
// 		Code:        "Item002",
// 		Title:       "Bola futebol",
// 		Description: "bola futebol",
// 		Price:       20,
// 		Stock:       30,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	},
// }

// type itemRepositoryImpl struct {
// }

// func NewRepository() ports.ItemRepository {
// 	return itemRepositoryImpl{}
// }

// func (s itemRepositoryImpl) GetAllItems() []domain.Item {
// 	return items
// }

// func (s itemRepositoryImpl) AddItem(item domain.Item) (*domain.Item, error) {
// 	for _, i := range items {
// 		if i.Code == item.Code {
// 			return nil, fmt.Errorf("duplicated entry")
// 		}
// 	}

// 	item.ID = len(items) + 1
// 	item.CreatedAt = time.Now()
// 	item.UpdatedAt = item.CreatedAt

// 	items = append(items, item)

// 	return &item, nil
// }

// func (s itemRepositoryImpl) ReadItem(id int) *domain.Item {
// 	for _, item := range items {
// 		if id == item.ID {
// 			return &item
// 		}
// 	}

// 	return nil
// }

// func (s itemRepositoryImpl) UpdateItem(id int, itemNew domain.Item) *domain.Item {
// 	for i, v := range items {
// 		if id == v.ID {
// 			itemNew.ID = v.ID
// 			itemNew.CreatedAt = v.CreatedAt
// 			itemNew.UpdatedAt = time.Now()
// 			items[i] = itemNew

// 			return &itemNew
// 		}
// 	}

// 	return nil
// }
