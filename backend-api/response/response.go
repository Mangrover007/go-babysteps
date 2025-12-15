package response

import (
	"net/http"
	"encoding/json"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
)

func BadRequest[T models.ResponseType] (w http.ResponseWriter, message string) error {
	w.WriteHeader(400)
	w.Header().Set("Content-Type", "application/json")

	res := models.ResponseBody[T]{
		Status: 400,
		Error: message,
		Data: nil,
	}
	
	err := json.NewEncoder(w).Encode(res)
	return err
}

func Success[T models.ResponseType] (w http.ResponseWriter, result T) error {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	res := models.ResponseBody[T]{
		Status: 200,
		Data: &(models.ResponseData[T]{
			Result: result,
		}),
	}

	err := json.NewEncoder(w).Encode(res)
	return err
}

func NotFound[T models.ResponseType](w http.ResponseWriter, message string) error {
	w.WriteHeader(http.StatusNotFound)
	
	res := models.ResponseBody[T]{
		Status: http.StatusNotFound,
		Error: message,
		Data: nil,
	}

	err := json.NewEncoder(w).Encode(res)
	return err
}

func NotAuthorized[T models.ResponseType](w http.ResponseWriter, message string) error {
	w.WriteHeader(http.StatusUnauthorized)

	res := models.ResponseBody[T]{
		Status: http.StatusUnauthorized,
		Error: message,
		Data: nil,
	}

	err := json.NewEncoder(w).Encode(res)
	return err
}

// func InternalServerError(w http.ResponseWriter) {
// 	w.WriteHeader(500)
// 	w.Write([]byte("Something went wrong..."))
// }

