package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createItemForm))
	mux.Post("/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createItem))

	mux.Get("/items", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showItems))
	mux.Get("/items/:id", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showItem))
	mux.Post("/items/update", http.HandlerFunc(app.updateItem))
	mux.Post("/items/delete", http.HandlerFunc(app.deleteItem))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	mux.Post("/cart/add", alice.New(app.session.Enable).ThenFunc(app.addCart))
	mux.Get("/cart", alice.New(app.session.Enable).ThenFunc(app.showCart))
	mux.Post("/cart/delete", alice.New(app.session.Enable).ThenFunc(app.deleteItemFromCart))
	mux.Get("/cart/buy", alice.New(app.session.Enable).ThenFunc(app.buyForm))
	mux.Post("/cart/buy", alice.New(app.session.Enable).ThenFunc(app.buy))

	mux.Get("/orders", alice.New(app.session.Enable).ThenFunc(app.showOrders))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
