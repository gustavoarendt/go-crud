package controllers

import (
	"go-crud/models"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceAsFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		amountAsInt, err := strconv.Atoi(amount)
		if err != nil {
			panic(err.Error())
		}

		models.InsertProduct(name, description, priceAsFloat, amountAsInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.EditProduct(id)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceAsFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}

		amountAsInt, err := strconv.Atoi(amount)
		if err != nil {
			panic(err.Error())
		}

		models.UpdateProduct(id, name, description, priceAsFloat, amountAsInt)
	}
	http.Redirect(w, r, "/", 301)
}
