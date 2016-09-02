package main

import (
	"net/http"

	"github.com/maximelamure/training/server/controllers"
)

func main() {

	ctr := controllers.NewProductsController()
	http.HandleFunc("/products", ctr.Index)
	http.ListenAndServe("localhost:8080", nil)

}
