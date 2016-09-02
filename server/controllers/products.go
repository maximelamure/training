package controllers

import (
	"encoding/json"
	"net/http"
)

type ProductsController struct {
}

func NewProductsController() *ProductsController {
	return &ProductsController{}
}

func (p *ProductsController) Index(resp http.ResponseWriter, r *http.Request) {
	products, err := GetProducts()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(resp).Encode(products)
}
