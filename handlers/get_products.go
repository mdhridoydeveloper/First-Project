package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func Getproduct(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.Productlist, 200)
}
