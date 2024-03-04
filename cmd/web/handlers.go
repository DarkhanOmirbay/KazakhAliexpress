package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "HTML ERROR", 500)
		return
	}
	ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
func (app *Application) createItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	name := r.Form.Get("iname")
	item_type := r.Form.Get("titem")
	priceStr := r.Form.Get("price")
	imgurl := r.Form.Get("img")
	quantityStr := r.Form.Get("qu")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "Error converting price to integer", http.StatusBadRequest)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(w, "Error converting quantity to integer", http.StatusBadRequest)
		return
	}

	app.ItemModel.Insert(app.DB, name, item_type, imgurl, price, quantity)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app *Application) showItems(w http.ResponseWriter, r *http.Request) {
	items, err := app.ItemModel.Read(app.DB)
	if err != nil {
		http.Error(w, "Error reading items from the database", http.StatusInternalServerError)
		return
	}

	ts, err := template.ParseFiles("./ui/html/items.page.tmpl")
	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, items)
	if err != nil {
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
		return
	}
}
func (app *Application) showItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		log.Fatal(err)
	}
	i, err := app.ItemModel.GetItem(app.DB, id)
	if err != nil {
		log.Fatal(err)
	}
	ts, err := template.ParseFiles("./ui/html/item.page.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = ts.Execute(w, i)
	if err != nil {
		log.Fatal(err)
	}

}
func (app *Application) updateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
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
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted value:", id)
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted value:", price)
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted value:", quantity)
	}
	err = app.ItemModel.Update(app.DB, name, item_type, imgurl, id, price, quantity)
	if err != nil {
		http.Error(w, "db error update", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/items", http.StatusSeeOther)
}
func (app *Application) deleteItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	idStr := r.Form.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Converted value:", id)
	}
	err = app.ItemModel.Delete(app.DB, id)
	if err != nil {
		http.Error(w, "db delete error", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/items", http.StatusSeeOther)

}
