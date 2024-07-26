package main

import (
	"log"

	"github.com/melisource/fury_go-platform/pkg/fury"
	"github.com/osalomon89/go-basics/internal/adapters/handler"
	mysqlrepo "github.com/osalomon89/go-basics/internal/adapters/repository/mysql"
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

	// esCient, err := es8.NewDefaultClient()
	// if err != nil {
	// 	return err
	// }

	// repo := ds.NewEsRepository(esCient)
	// if err := repo.CreateIndex("items"); err != nil {
	// 	log.Fatalln(err)
	// }

	conn, err := mysqlrepo.GetConnectionDB()
	if err != nil {
		return err
	}

	repo := mysqlrepo.NewMySQLRepository(conn)
	service := service.NewService(repo) //itemServiceImpl

	h := handler.NewHandler(service) //ItemService

	app.Get("/hello", h.HelloHandler)
	app.Get("/items", h.ReadItem)
	app.Post("/items", h.CreateItem)
	app.Get("/items/{id}", h.ReadItemId)
	app.Put("/items/{id}", h.UpdateItem)
	//adicionar delete endpoint

	return app.Run()
}
