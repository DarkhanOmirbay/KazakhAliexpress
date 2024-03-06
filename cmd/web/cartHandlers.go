package main

import (
	"KazakhAliexpress/SE2220/pkg/forms"
	"KazakhAliexpress/SE2220/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) addCart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	ItemIdStr := r.Form.Get("ItemId")
	ItemId, err := strconv.Atoi(ItemIdStr)
	if err != nil {
		app.errorlog.Println("Error", err)
		http.Error(w, "Invalid ItemId parameter", http.StatusBadRequest)
		return
	}
	UserId := app.session.GetInt(r, "authenticatedUserID")

	app.errorlog.Println("Before AddCart:", ItemId, UserId)

	app.carts.AddCart(ItemId, UserId)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *application) showCart(w http.ResponseWriter, r *http.Request) {

	userId := app.session.GetInt(r, "authenticatedUserID")

	carts, err := app.carts.GetCartsById(userId)
	if err != nil {
		app.errorlog.Println(err)
	}

	var items []*models.Item

	for _, cart := range carts {
		item, err := app.items.GetItem(cart.ItemId)
		if err != nil {
			app.errorlog.Println(err)
		}

		items = append(items, item)
	}

	data := &templateData{
		Items: items,
	}

	app.render(w, r, "cart.page.tmpl", data)
}
func (app *application) deleteItemFromCart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}
	ItemIdStr := r.Form.Get("ItemId")
	ItemId, err := strconv.Atoi(ItemIdStr)
	if err != nil {
		app.errorlog.Println("Error", err)
		http.Error(w, "Invalid ItemId parameter", http.StatusBadRequest)
		return
	}
	app.carts.DeleteItem(ItemId)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}
func (app *application) buyForm(w http.ResponseWriter, r *http.Request) {
	userId := app.session.GetInt(r, "authenticatedUserID")

	// Получаем список карточек пользователя
	carts, err := app.carts.GetCartsById(userId)
	if err != nil {
		app.errorlog.Println(err)
	}

	var items []*models.Item

	for _, cart := range carts {
		item, err := app.items.GetItem(cart.ItemId)
		if err != nil {
			app.errorlog.Println(err)
		}

		items = append(items, item)
	}

	app.render(w, r, "buy.page.tmpl", &templateData{Form: forms.New(nil),
		Items: items})
}
func (app *application) buy(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("address", "cardNumber", "expirationDate", "cvv")
	form.MaxLength("address", 50)
	form.MaxLength("cardNumber", 16)
	form.MaxLength("expirationDate", 5)
	form.MaxLength("cvv", 3)
	form.MinLength("address", 10)
	form.MinLength("cardNumber", 16)
	form.MinLength("expirationDate", 5)
	form.MinLength("cvv", 3)

	if !form.Valid() {
		app.render(w, r, "buy.page.tmpl", &templateData{Form: form})
		return
	}
	userId := app.session.GetInt(r, "authenticatedUserID")
	msg := "Your orders will be delivery within 10 days to the address:" + r.Form.Get("address")
	app.orders.Insert(userId, form.Get("address"), msg)
	app.users.Update(userId, form.Get("cardNumber"), form.Get("expirationDate"), form.Get("cvv"))
	http.Redirect(w, r, "/orders", http.StatusSeeOther)

}
func (app *application) showOrders(w http.ResponseWriter, r *http.Request) {
	userId := app.session.GetInt(r, "authenticatedUserID")
	order, err := app.orders.GetOrdersByUserId(userId)
	if err != nil {
		app.errorlog.Println(err)
	}
	carts, err := app.carts.GetCartsById(order.UserId)
	if err != nil {
		app.errorlog.Println(err)
	}
	var items []*models.Item

	for _, cart := range carts {
		item, err := app.items.GetItem(cart.ItemId)
		if err != nil {
			app.errorlog.Println(err)
		}

		items = append(items, item)
	}
	data := &templateData{Order: order, Carts: carts, Items: items}
	//app.carts.DeleteItemUseUserId(userId)
	app.render(w, r, "orders.page.tmpl", data)
}
