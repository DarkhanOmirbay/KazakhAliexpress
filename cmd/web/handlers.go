package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
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

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		app.serverError(w, err)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.items.Insert(name, item_type, imgurl, price, quantity)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *application) showItems(w http.ResponseWriter, r *http.Request) {
	items, err := app.items.Read()
	if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/items.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, items)
	if err != nil {
		app.serverError(w, err)
		return
	}
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
	files := []string{
		"./ui/html/item.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}
	err = ts.Execute(w, i)
	if err != nil {
		log.Fatal(err)
	}

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

	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.errorlog.Println("Error:", err)
	} else {
		app.errorlog.Println("Converted value:", id)
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		app.errorlog.Println("Error:", err)
	} else {
		app.errorlog.Println("Converted value:", price)
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		app.errorlog.Println("Error:", err)
	} else {
		app.errorlog.Println("Converted value:", quantity)
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
