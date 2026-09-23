package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	manager.Use(middleware.Hudai, middleware.Logger)
	mux := http.NewServeMux()
	globalRouter := global_router.GlobalRouter(mux)
	InitRoutes(mux, manager)

	fmt.Println("your sarver running on:8080")

	err := http.ListenAndServe(":8000", globalRouter)
	if err != nil {
		fmt.Println("Error your sarver", err)
	}

}
