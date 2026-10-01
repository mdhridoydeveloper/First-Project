package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var Newproduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&Newproduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me  vaild json", 400)
		return
	}
	Newproduct.ID = pId
	database.Update(Newproduct)
	util.SendData(w, "Successfully updated product", 201)
}
