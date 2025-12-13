package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"
)

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.BadRequest[int](w, "Method must be a POST request")
		return
	}

	var body models.RequestBodyOther
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.BadRequest[int](w, err.Error())
		return
	}

	response.Success[int64](w, int64(body.Num1) * int64(body.Num2))
}

