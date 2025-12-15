package auth

import (
	"net/http"
	"encoding/json"
	"errors"

	"github.com/Mangrover007/go-babysteps/backend-api/internals/response"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/db"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var body models.RequestForLogin
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		var TypeError *json.UnmarshalTypeError
		if errors.As(err, &TypeError) {
			response.BadRequest(w, "type mismatch")
			return
		}
		response.InternalServerError(w)
		return
	}

	// input validation
	if body.Email == "" || body.Password == "" {
		response.BadRequest(w, "email and password must be non-empty")
		return
	}

	var users *(map[string]string) = &(db.Users)

	pass, ok := (*users)[body.Email];
	if !ok {
		response.NotFound(w, "user not found in the database register first dumbass")
		return
	}
	if pass != body.Password {
		response.NotAuthorized(w, "password did not match")
		return
	}

	cookie := http.Cookie{
		Name: "token",
		Value: "BADAPPLE",
		Path: "/",
	}
	http.SetCookie(w, &cookie)
	response.Success[string](w, "log in successfull")
	return
}

