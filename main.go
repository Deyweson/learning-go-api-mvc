package main

import (
	"html/template"
	"net/http"

	"github.com/deyweson/learning-go-api-mvc/models"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)

	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
