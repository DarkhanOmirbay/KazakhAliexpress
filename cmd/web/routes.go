package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Post("/create", dynamicMiddleware.ThenFunc(app.createItem))

	mux.Get("/items", dynamicMiddleware.ThenFunc(app.showItems))
	mux.Get("/items/:id", dynamicMiddleware.ThenFunc(app.showItem))
	mux.Post("/items/update", dynamicMiddleware.ThenFunc(app.updateItem))
	mux.Post("/items/delete", dynamicMiddleware.ThenFunc(app.deleteItem))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
