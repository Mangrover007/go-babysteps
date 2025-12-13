package models

type RequestBodyOther struct {
	Num1 int `json:"number1"`
	Num2 int `json:"number2"`
}

type RequestBodyDivide struct {
	Divisor 	float64 `json:"divisor"`
	Dividend 	float64 `json:"dividend"`
}

type RequestBodySum struct {
	Items []int `json:"items"`
}



type ResultType interface {
	int | int64 | float64
}

type ResponseData[T ResultType] struct {
	Result T `json:"result"`
}

type ResponseBody[T ResultType] struct {
	Data 	*(ResponseData[T]) 	`json:"data,omitempty"`
	Status 	int 				`json:"status"`
	Error 	string 				`json:"error,omitempty"`
}

