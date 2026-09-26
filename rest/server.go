package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Corss,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	InitRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Sarver running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error your sarver", err)
		os.Exit(1)
	}

}
