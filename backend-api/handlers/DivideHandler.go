package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.BadRequest[int](w, "Method must be a POST request")
		return
	}

	var body models.RequestBodyDivide
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.BadRequest[int](w, err.Error())
		return
	}

	if body.Divisor == float64(0) {
		response.BadRequest[int](w, "Cannot divide by 0")
		return
	}

	response.Success[float64](w, body.Dividend/body.Divisor)
}

