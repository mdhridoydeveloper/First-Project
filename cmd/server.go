package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Hudai, middleware.Logger)
	mux := http.NewServeMux()

	mux.Handle(
		"GET /Hidoy",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	mux.Handle(
		"GET /habib", manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.Getproduct),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.Createproduct),
		),
	)
	mux.Handle(
		"GET /product/{productID}",
		manager.With(
			http.HandlerFunc(handlers.GetproductByID),
		),
	)
	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("your sarver running on:8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error your sarver", err)
	}

}
