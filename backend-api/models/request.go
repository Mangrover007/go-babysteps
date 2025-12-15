package models

// skip the interface idk what to do with it lol
// type Request interface {
// 	RequestForDivision | RequestForOtherOps | RequestForLogin | RequestForRegister
// }

type RequestForDivision struct {
	Divisor		float64 `json:"divisor,omitempty"`
	Dividend	float64 `json:"dividend,omitempty"`
}

type RequestForOtherOps struct {
	Number1 	int64 `json:"number1,omitempty"`
	Number2		int64 `json:"number2,omitempty"`
}

type RequestForSum struct {
	Items		[]int64 `json:"items,omitempty"`
}

type RequestForLogin struct {
	Email		string `json:"email,omitempty"`
	Password	string `json:"password,omitempty"`
}

type RequestForRegister struct {
	Email		string `json:"email,omitempty"`
	Password	string `json:"password,omitempty"`
	Username	string `json:"username,omitempty"`
}

