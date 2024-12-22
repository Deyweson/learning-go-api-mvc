package controllers

import (
	"net/http"
	"text/template"

	"github.com/deyweson/learning-go-api-mvc/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
