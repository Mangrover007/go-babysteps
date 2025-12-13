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
)

func main() {
	http.Handle("/add", new(handlers.AddHandler))
	http.HandleFunc("/subtract", handlers.SubtractHandler)
	http.HandleFunc("/multiply", handlers.MultiplyHandler)
	http.HandleFunc("/divide", handlers.DivideHandler)
	http.HandleFunc("/sum", handlers.SumHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

