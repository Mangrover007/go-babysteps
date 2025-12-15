package main

import (
	"net/http"
	"fmt"

	"github.com/Mangrover007/go-babysteps/backend-api/internals/routers"
	"github.com/Mangrover007/go-babysteps/backend-api/internals/middlewares"
)

const PORT int = 3000

func main() {
	Router := http.NewServeMux()
	Router.Handle("/auth/", http.StripPrefix("/auth", routers.AuthRouter()))
	Router.Handle("/",      middlewares.Protect(routers.ProtectedRouter()))

	Server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: middlewares.Logger(Router),
	}

	fmt.Printf("Server listening on port %d\n", PORT)
	err := Server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

