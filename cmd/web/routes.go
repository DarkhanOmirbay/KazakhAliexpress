package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *Application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Post("/create", http.HandlerFunc(app.createItem))

	mux.Get("/items", http.HandlerFunc(app.showItems))
	mux.Get("/items/:id", http.HandlerFunc(app.showItem))
	mux.Post("/items/update", http.HandlerFunc(app.updateItem))
	mux.Post("/items/delete", http.HandlerFunc(app.deleteItem))
	return mux
}
