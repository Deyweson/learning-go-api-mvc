package routes

import (
	"net/http"

	"github.com/deyweson/learning-go-api-mvc/controllers"
)

func InicializarRotas() {
	http.HandleFunc("/", controllers.Index)
}
