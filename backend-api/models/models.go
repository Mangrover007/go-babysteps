package models

import "fmt"

type RequestBodyOther struct {
	Num1 int `json:"number1,omitempty"`
	Num2 int `json:"number2,omitempty"`
}

type RequestBodyDivide struct {
	Divisor 	float64 `json:"divisor,omitempty"`
	Dividend 	float64 `json:"dividend,omitempty"`
}

type RequestBodySum struct {
	Items []int `json:"items,omitempty"`
}

type RequestLogin struct {
	Email 		string `json:email,omitempty`
	Password 	string `json:password,omitempty`
}



type ResponseType interface {
	int | int64 | float64 | string
}

type ResponseData[T ResponseType] struct {
	Result T `json:"result"`
}

type ResponseBody[T ResponseType] struct {
	Data 	*(ResponseData[T]) 	`json:"data,omitempty"`
	Status 	int 				`json:"status"`
	Error 	string 				`json:"error,omitempty"`
}


// mock db lol
type DB struct {
    Users         map[string]string
    LoggedInUsers map[string]bool
}

var MockDB = DB{
    Users: map[string]string{
        "john":   "doe",
        "selena": "gomez",
        "opm":    "trash", 
    },
    LoggedInUsers: map[string]bool{}, 
}

func main() {
    fmt.Printf("Mock DB initialized with %d users.\n", len(MockDB.Users))
    if pwd, ok := MockDB.Users["john"]; ok {
        fmt.Println("John's password:", pwd)
    }
}

