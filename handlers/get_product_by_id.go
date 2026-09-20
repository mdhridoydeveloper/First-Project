package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetproductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productID")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}
	for _, product := range database.Productlist {
		if product.ID == pId {
			util.SendData(w, product, 200)
		}
	}
	util.SendData(w, "`No Show any data", 404)
}
