package middleware

import (
	"fmt"
	"net/http"
	"github.com/Mangrover007/go-babysteps/backend-api/response"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("now the codebase is schmoozing")
		token, err := r.Cookie("token")
		if err != nil {
			response.NotAuthorized[int](w, "\"token\" is missing from cookies ping /login with correct credentials to get a cookie")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

