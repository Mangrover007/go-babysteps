package prot

import (
	"net/http"
	"encoding/json"
	"errors"

	"github.com/Mangrover007/go-babysteps/backend-api/internals/response"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {

	var body models.RequestForDivision
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		var TypeError *json.UnmarshalTypeError
		if errors.As(err, &TypeError) {
			response.BadRequest(w, "type mistmatch expected float64")
			return
		}

		response.InternalServerError(w)
		return
	}

	if body.Divisor == float64(0) {
		response.BadRequest(w, "divide by 0 is not allowed send different divisor")
		return
	}

	response.Success[float64](w, body.Dividend / body.Divisor)
	return 
}

