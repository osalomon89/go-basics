package main

import (
	"log"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/melisource/fury_go-platform/pkg/fury"
	"github.com/osalomon89/go-basics/internal/adapters/handler"
	repository "github.com/osalomon89/go-basics/internal/adapters/repository/ds"
	"github.com/osalomon89/go-basics/internal/core/service"
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

	esCient, err := es8.NewDefaultClient()
	if err != nil {
		return err
	}

	repo := repository.NewEsRepository(esCient)
	if err := repo.CreateIndex("items"); err != nil {
		log.Fatalln(err)
	}

	// repo := mysqlrepo.NewMySQLRepository()
	service := service.NewService(repo) //itemServiceImpl

	h := handler.NewHandler(service) //ItemService

	app.Get("/hello", h.HelloHandler)
	app.Get("/items", h.ReadItem)
	app.Post("/items", h.CreateItem)
	app.Get("/items/{id}", h.ReadItemId)
	app.Put("/items/{id}", h.UpdateItem)

	return app.Run()
}
