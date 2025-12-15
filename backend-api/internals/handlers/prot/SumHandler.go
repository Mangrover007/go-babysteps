package prot

import (
	"net/http"
	"encoding/json"
	"errors"

	"github.com/Mangrover007/go-babysteps/backend-api/internals/response"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {

	var body models.RequestForSum
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		var TypeError *json.UnmarshalTypeError
		if errors.As(err, &TypeError) {
			response.BadRequest(w, "type mistmatch expected items array of int64")
			return
		}

		response.InternalServerError(w)
		return
	}

	var sum int64 = 0
	for _, item := range body.Items {
		sum += item
	}

	response.Success[int64](w, sum)
	return 
}

