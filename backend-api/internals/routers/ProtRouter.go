package routers

import (
	"net/http"
	"github.com/Mangrover007/go-babysteps/backend-api/internals/handlers/prot"
)

func ProtectedRouter() http.Handler {
	
	router := http.NewServeMux()

	// register all the fucking routers
	router.HandleFunc("POST /add", 		prot.AddHandler)
	router.HandleFunc("POST /divide", 	prot.DivideHandler)
	router.HandleFunc("POST /multiply", prot.MultiplyHandler)
	router.HandleFunc("POST /subtract", prot.SubtractHandler)
	router.HandleFunc("POST /sum", 		prot.SumHandler)

	return router
}

