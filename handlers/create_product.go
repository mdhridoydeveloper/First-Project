package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func Createproduct(w http.ResponseWriter, r *http.Request) {

	var Newproduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&Newproduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me  vaild json", 400)
		return
	}
	Newproduct.ID = len(database.Productlist) + 1
	database.Productlist = append(database.Productlist, Newproduct)

	util.SendData(w, Newproduct, 200)
}
