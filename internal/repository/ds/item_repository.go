package ds

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/osalomon89/go-basics/internal/core/domain"
)

type elasticSearch struct {
	client *es8.Client
	index  string
}

func NewEsRepository(client *es8.Client) *elasticSearch {
	return &elasticSearch{
		client: client,
	}
}

func (r *elasticSearch) GetAllItems() []domain.Item {
	return nil
}

func (r *elasticSearch) AddItem(ctx context.Context, item domain.Item) (*domain.Item, error) {
	now := time.Now().UTC()
	item.CreatedAt = &now
	item.UpdatedAt = &now

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(item); err != nil {
		return nil, fmt.Errorf("error encoding document: %s", err)
	}

	req := esapi.IndexRequest{
		Index: r.index,
		Body:  &buf,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, fmt.Errorf("insert: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 409 {
		return nil, errors.New("conflict")
	}

	if res.IsError() {
		return nil, fmt.Errorf("insert: response: %s", res.String())
	}

	return &item, nil
}

func (r *elasticSearch) Update(ctx context.Context, itemNew domain.Item) (*domain.Item, error) {
	now := time.Now().UTC()
	itemNew.UpdatedAt = &now
	bdy, err := json.Marshal(itemNew)
	if err != nil {
		return nil, fmt.Errorf("update: marshall: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      r.index,
		DocumentID: itemNew.ID,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, fmt.Errorf("update: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}

	if res.IsError() {
		return nil, fmt.Errorf("update: response: %s", res.String())
	}

	return &itemNew, nil
}

func (r *elasticSearch) Delete(ctx context.Context, id string) error {
	req := esapi.DeleteRequest{
		Index:      r.index,
		DocumentID: id,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("delete: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return fmt.Errorf("not found")
	}

	if res.IsError() {
		return fmt.Errorf("delete: response: %s", res.String())
	}

	return nil
}

func (r *elasticSearch) ReadItem(ctx context.Context, id string) (*domain.Item, error) {
	req := esapi.GetRequest{
		Index:      r.index,
		DocumentID: id,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return nil, fmt.Errorf("find one: request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}

	if res.IsError() {
		return nil, fmt.Errorf("find one: response: %s", res.String())
	}

	var (
		item domain.Item
		body document
	)
	body.Source = &item

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("find one: decode: %w", err)
	}

	return &item, nil
}

// document represents a single document in Get API response body.
type document struct {
	Source interface{} `json:"_source"`
}
