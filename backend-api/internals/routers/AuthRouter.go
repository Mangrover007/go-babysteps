package routers

import (
	"net/http"
	"github.com/Mangrover007/go-babysteps/backend-api/internals/handlers/auth"
)

func AuthRouter() http.Handler {
	
	router := http.NewServeMux()
	router.HandleFunc("POST /login",    auth.LoginHandler)
	router.HandleFunc("POST /register", auth.RegisterHandler)

	return router
}

