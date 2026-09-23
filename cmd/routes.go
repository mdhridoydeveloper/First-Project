package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
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
}
