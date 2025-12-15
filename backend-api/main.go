package main

// paths - 
// /add      - add two numbers
// /subtract - subtract two numbers
// /multiply - multiply two numbers
// /divide   - divide two numbers
// /sum      - add all numbers in array

// all numbers are int type

// req body - JSON
// res body - JSON

// req body for /add, /subtract, /multiply
// { number1: <>, number2: <> }

// for /divide
// { dividend: <>, divisor: <> }

// for /sum
// { items: []int }

// res body for all paths
// { result: <>int }

import (
	// "fmt"
	"net/http"
	"github.com/Mangrover007/go-babysteps/backend-api/handlers"
	"github.com/Mangrover007/go-babysteps/backend-api/middleware"
)

func main() {
	
	app := http.NewServeMux()

	// all unauthorized routes to this guy
	app.HandleFunc("/login", handlers.Login)

	// authorized app, register authorized routes to this guy
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /add", 		handlers.AddHandler)
	authRouter.HandleFunc("POST /divide", 	handlers.DivideHandler)
	authRouter.HandleFunc("POST /multiply", handlers.MultiplyHandler)
	authRouter.HandleFunc("POST /subtract", handlers.SubtractHandler)
	authRouter.HandleFunc("POST /sum", 		handlers.SumHandler)

	// add app to central router
	app.Handle("/", middleware.Authorize(authRouter))

	server := http.Server{
		Addr:		":3000",
		Handler:	app,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

