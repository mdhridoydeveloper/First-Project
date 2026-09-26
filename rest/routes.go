package rest

import (
	"ecommerce/rest/handlers"

	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.Getproducts),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.Createproducts),
		),
	)
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.Getproduct),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProducts),
		),
	)

}
