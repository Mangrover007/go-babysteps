package prot

import (
	"net/http"
	"encoding/json"
	"errors"

	"github.com/Mangrover007/go-babysteps/backend-api/internals/response"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {

	var body models.RequestForOtherOps
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		var TypeError *json.UnmarshalTypeError
		if errors.As(err, &TypeError) {
			response.BadRequest(w, "type mistmatch expected int64")
			return
		}

		response.InternalServerError(w)
		return
	}

	response.Success[int64](w, body.Number1 + body.Number2)
	return 
}

