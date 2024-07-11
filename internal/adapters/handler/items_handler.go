package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/melisource/fury_go-core/pkg/web"
	"github.com/osalomon89/go-basics/internal/core/domain"
	"github.com/osalomon89/go-basics/internal/core/ports"
)

type responseError struct {
	Message    string
	StatusCode int
}

type handler struct {
	itemService ports.ItemService
}

func NewHandler(itemService ports.ItemService) *handler {
	return &handler{
		itemService: itemService,
	}
}

func (h *handler) CreateItem(w http.ResponseWriter, r *http.Request) error {
	var newItem domain.Item

	if err := web.DecodeJSON(r, &newItem); err != nil {
		log.Printf("Failed to decode JSON: %v", err)
		return web.EncodeJSON(w, responseError{Message: "error decoding json body", StatusCode: http.StatusBadRequest}, http.StatusBadRequest)
	}

	item, err := h.itemService.AddItem(r.Context(), newItem)
	if err != nil {
		log.Printf("error inserting item: %v", err)
		return web.EncodeJSON(w, responseError{Message: "error inserting item", StatusCode: http.StatusInternalServerError}, http.StatusInternalServerError)
	}

	return web.EncodeJSON(w, item, http.StatusCreated)
}

func (h *handler) ReadItemId(w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")

	item := h.itemService.ReadItem(r.Context(), id)
	if item != nil {
		return web.EncodeJSON(w, item, http.StatusOK)
	}

	return web.EncodeJSON(w, responseError{Message: "id not found", StatusCode: http.StatusNotFound}, http.StatusNotFound)
}

func (h *handler) ReadItem(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	limitParam := r.URL.Query().Get("limit")
	cursorParam := r.URL.Query().Get("cursor")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	var searchAfter []interface{}
	if cursorParam != "" {
		searchAfter = parseCursor(cursorParam)
	}

	items, newCursor, err := h.itemService.GetAllItems(ctx, limit, searchAfter)
	if err != nil {
		return web.EncodeJSON(w, responseError{Message: "error getting items", StatusCode: http.StatusInternalServerError}, http.StatusInternalServerError)
	}

	response := struct {
		Items  []domain.Item `json:"items"`
		Cursor string        `json:"cursor,omitempty"`
	}{
		Items: items,
	}

	if newCursor != nil {
		response.Cursor = encodeCursor(newCursor)
	}

	return web.EncodeJSON(w, response, http.StatusOK)
}

func parseCursor(cursor string) []interface{} {
	parts := strings.Split(cursor, ",")
	parsed := make([]interface{}, len(parts))
	for i, part := range parts {
		if num, err := strconv.ParseFloat(part, 64); err == nil {
			parsed[i] = num
		} else {
			parsed[i] = part
		}
	}
	return parsed
}

func encodeCursor(cursor []interface{}) string {
	parts := make([]string, len(cursor))
	for i, part := range cursor {
		switch v := part.(type) {
		case float64:
			// Convertimos el número a un string sin formato exponencial
			parts[i] = strconv.FormatFloat(v, 'f', -1, 64)
		default:
			parts[i] = fmt.Sprintf("%v", part)
		}
	}
	return strings.Join(parts, ",")
}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request) error {
	var existItem domain.Item

	if err := web.DecodeJSON(r, &existItem); err != nil {
		log.Printf("Failed to decode JSON: %v", err)
		return web.EncodeJSON(w, responseError{Message: "error decoding json body", StatusCode: http.StatusBadRequest}, http.StatusBadRequest)
	}

	id := web.Param(r, "id")
	existItem.ID = id

	result := h.itemService.UpdateItem(r.Context(), existItem)

	if result != nil {
		return web.EncodeJSON(w, result, http.StatusOK)
	}

	return web.EncodeJSON(w, responseError{Message: "id not found", StatusCode: http.StatusNotFound}, http.StatusNotFound)
}

func (h *handler) HelloHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, fmt.Sprintf("%s, world!", r.URL.Path[1:]), http.StatusOK)
}
