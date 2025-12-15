package handlers

import (
	"errors"
	"net/http"
	"encoding/json"

	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.BadRequest[int](w, "method must be a POST request")
		return
	}

	var credentials models.RequestLogin
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		var typeError *json.UnmarshalTypeError

		if errors.As(err, &typeError) {
			response.BadRequest[int](w, "type mismatch")
			return
		}

		return
	}

	val, ok := models.MockDB.Users[credentials.Email]
	if !ok {
		response.NotFound[int](w, "email not found. register at /register")
		return
	}

	if val != credentials.Password {
		response.NotAuthorized[int](w, "incorrect password fuck off")
		return
	}

	// simulating log in i am NOT adding a database NO
	models.MockDB.LoggedInUsers[credentials.Email] = true
	tokCookie := http.Cookie{
		Name: "token",
		Value: "literally what the fuck?"
	}

	w.SetCookie(w, &tokCookie)
	response.Success[string](w, "login successfull")
}

