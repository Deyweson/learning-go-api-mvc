package main

import (
	"net/http"

	"github.com/deyweson/learning-go-api-mvc/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.InicializarRotas()
	http.ListenAndServe(":8000", nil)
}
