package middlewares

import (
	"net/http"

	"github.com/Mangrover007/go-babysteps/backend-api/internals/response"
)

func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		// verify token login im just gonna check if the cookie is there or not
		// and if its value is "BADAPPLE" or not

		if err != nil {
			response.NotAuthorized(w, "get a token so login first")
			return
		}

		if token.Value != "BADAPPLE" {
			response.NotAuthorized(w, "invalid token")
			return
		}

		// forward the request to the next handler, so call its ServeHTTP method
		next.ServeHTTP(w, r)
	})
}

