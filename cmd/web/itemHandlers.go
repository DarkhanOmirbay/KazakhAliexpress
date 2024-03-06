package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	items, err := app.items.Read()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Items: items}

	app.render(w, r, "home.page.tmpl", data)
}
func (app *application) createItemForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{})
}
func (app *application) createItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	name := r.Form.Get("iname")
	item_type := r.Form.Get("titem")
	priceStr := r.Form.Get("price")
	imgurl := r.Form.Get("img")
	quantityStr := r.Form.Get("qu")

	errors := make(map[string]string)
	if strings.TrimSpace(name) == "" {
		errors["name"] = "This field cant be blank"
	} else if utf8.RuneCountInString(name) > 100 {
		errors["name"] = "This field is too long (maximum is 100 characters)"
	}
	if strings.TrimSpace(item_type) == "" {
		errors["item_type"] = "This field cant be blank"
	} else if utf8.RuneCountInString(item_type) > 100 {
		errors["name"] = "This field is too long (maximum is 100 characters)"
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if strings.TrimSpace(priceStr) == "" {
		errors["priceStr"] = "This field can not be blank"
	} else if price <= 0 {
		errors["priceStr"] = "This field is invalid"
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if strings.TrimSpace(quantityStr) == "" {
		errors["quantityStr"] = "This field can not be blank"
	} else if quantity <= 0 {
		errors["quantityStr"] = "This field is invalid"
	}
	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	app.items.Insert(name, item_type, imgurl, price, quantity)
	app.session.Put(r, "flash", "Snippet successfully created!")

	http.Redirect(w, r, "/items", http.StatusSeeOther)
}
func (app *application) showItems(w http.ResponseWriter, r *http.Request) {

	items, err := app.items.Read()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Items: items}
	app.render(w, r, "items.page.tmpl", data)
}
func (app *application) showItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.serverError(w, err)
		return
	}
	i, err := app.items.GetItem(id)
	if err != nil {
		log.Fatal(err)
	}
	data := &templateData{Item: i}

	app.render(w, r, "item.page.tmpl", data)

}
func (app *application) updateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}
	idStr := r.Form.Get("id")
	name := r.Form.Get("iname")
	item_type := r.Form.Get("titem")
	priceStr := r.Form.Get("price")
	imgurl := r.Form.Get("img")
	quantityStr := r.Form.Get("qu")

	errors := make(map[string]string)
	if strings.TrimSpace(name) == "" {
		errors["name"] = "This field cant be blank"
	} else if utf8.RuneCountInString(name) > 100 {
		errors["name"] = "This field is too long (maximum is 100 characters)"
	}
	if strings.TrimSpace(item_type) == "" {
		errors["item_type"] = "This field cant be blank"
	} else if utf8.RuneCountInString(item_type) > 100 {
		errors["name"] = "This field is too long (maximum is 100 characters)"
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if strings.TrimSpace(priceStr) == "" {
		errors["priceStr"] = "This field can not be blank"
	} else if price <= 0 {
		errors["priceStr"] = "This field is invalid"
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if strings.TrimSpace(quantityStr) == "" {
		errors["quantityStr"] = "This field can not be blank"
	} else if quantity <= 0 {
		errors["quantityStr"] = "This field is invalid"
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.errorlog.Println("Error:", err)
	} else {
		app.errorlog.Println("Converted value:", id)
	}

	if len(errors) > 0 {
		app.render(w, r, "item.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	err = app.items.Update(name, item_type, imgurl, id, price, quantity)
	if err != nil {
		http.Error(w, "db error update", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/items", http.StatusSeeOther)
}
func (app *application) deleteItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}
	idStr := r.Form.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.errorlog.Println("Error", err)
	} else {
		app.errorlog.Println("Converted value:", id)
	}
	err = app.items.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, "/items", http.StatusSeeOther)

}
