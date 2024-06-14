package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/melisource/fury_go-core/pkg/web"
	"github.com/melisource/fury_go-platform/pkg/fury"
	"github.com/osalomon89/go-basics/core/domain"
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

	app.Get("/hello", helloHandler)
	app.Post("/items", createItemHandler)

	return app.Run()
}

func helloHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, fmt.Sprintf("%s, world!", r.URL.Path[1:]), http.StatusOK)
}

type responseError struct {
	Message    string
	StatusCode int
}

func createItemHandler(w http.ResponseWriter, r *http.Request) error {
	log.Println("Entering ItemHandler: newItemHandler()")

	var newItem domain.Item

	if err := web.DecodeJSON(r, &newItem); err != nil {
		log.Printf("Failed to decode JSON: %v", err)
		return web.EncodeJSON(w, responseError{Message: "error decoding json body", StatusCode: http.StatusBadRequest}, http.StatusBadRequest)
	}

	log.Println("item created: ", newItem.Name)

	return web.EncodeJSON(w, newItem, http.StatusOK)
}
