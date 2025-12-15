package auth

import (
	"net/http"
	"encoding/json"
	"errors"

	"github.com/Mangrover007/go-babysteps/backend-api/internals/response"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/db"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	
	var body models.RequestForRegister
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		var TypeError *json.UnmarshalTypeError
		if errors.As(err, &TypeError) {
			response.BadRequest(w, "type mismtach")
			return
		}

		response.InternalServerError(w)
		return
	}

	// input validation yet again
	if body.Email == "" || body.Username == "" || body.Password == "" {
		response.BadRequest(w, "no field must be empty")
		return
	}

	users := &(db.Users)
	loggedInUsers := &(db.LoggedInUsers)

	(*users)[body.Email] = body.Password
	(*loggedInUsers)[body.Email] = true

	response.Success[string](w, "registration sucessfull. login to obtain a token")
	return
}

