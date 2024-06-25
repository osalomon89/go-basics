package main

import (
	"log"

	"github.com/melisource/fury_go-platform/pkg/fury"
	"github.com/osalomon89/go-basics/internal/handler"
	"github.com/osalomon89/go-basics/internal/repository"
	"github.com/osalomon89/go-basics/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}

	repo := repository.NewRepository()
	service := service.NewService(repo)
	h := handler.NewHandler(service)

	app.Get("/hello", h.HelloHandler)
	app.Get("/items", h.ReadItem)
	app.Post("/items", h.CreateItem)
	app.Get("/items/:id", h.ReadItemId)
	app.Put("/items/:id", h.UpdateItem)

	return app.Run()
}
