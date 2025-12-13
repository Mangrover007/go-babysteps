package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.BadRequest[int](w, "Method must be a POST request")
		return
	}

	var body models.RequestBodySum
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.BadRequest[int](w, err.Error())
		return
	}

	var sum int64 = 0
	for _, num := range body.Items {
		sum += int64(num)
	}

	response.Success[int64](w, sum)
}

