package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	item, err := h.itemService.AddItem(newItem)
	if err != nil {
		log.Printf("error inserting item: %v", err)
		return web.EncodeJSON(w, responseError{Message: "error inserting item", StatusCode: http.StatusInternalServerError}, http.StatusInternalServerError)
	}

	return web.EncodeJSON(w, item, http.StatusCreated)
}

func (h *handler) ReadItemId(w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return web.EncodeJSON(w, responseError{Message: "error in param", StatusCode: http.StatusBadRequest}, http.StatusBadRequest)
	}

	item := h.itemService.ReadItem(convertId)
	if item != nil {
		return web.EncodeJSON(w, item, http.StatusOK)
	}

	return web.EncodeJSON(w, responseError{Message: "id not found", StatusCode: http.StatusNotFound}, http.StatusNotFound)
}

func (h *handler) ReadItem(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, h.itemService.GetAllItems(), http.StatusOK)
}

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request) error {
	var existItem domain.Item

	if err := web.DecodeJSON(r, &existItem); err != nil {
		log.Printf("Failed to decode JSON: %v", err)
		return web.EncodeJSON(w, responseError{Message: "error decoding json body", StatusCode: http.StatusBadRequest}, http.StatusBadRequest)
	}

	id := web.Param(r, "id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return web.EncodeJSON(w, responseError{Message: "error in param", StatusCode: http.StatusBadRequest}, http.StatusBadRequest)
	}

	result := h.itemService.UpdateItem(convertId, existItem)

	if result != nil {
		return web.EncodeJSON(w, result, http.StatusOK)
	}

	return web.EncodeJSON(w, responseError{Message: "id not found", StatusCode: http.StatusNotFound}, http.StatusNotFound)
}

func (h *handler) HelloHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, fmt.Sprintf("%s, world!", r.URL.Path[1:]), http.StatusOK)
}
