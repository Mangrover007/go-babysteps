package response

import (
	"net/http"
	"encoding/json"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
)

func BadRequest[T models.ResultType] (w http.ResponseWriter, message string) {
	w.WriteHeader(400)
	w.Header().Set("Content-Type", "application/json")

	res := models.ResponseBody[T]{
		Status: 400,
		Error: message,
		Data: nil,
	}
	
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
}

func Success[T models.ResultType] (w http.ResponseWriter, result T) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	res := models.ResponseBody[T]{
		Status: 200,
		Data: &(models.ResponseData[T]{
			Result: result,
		}),
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
}

// func InternalServerError(w http.ResponseWriter) {
// 	w.WriteHeader(500)
// 	w.Write([]byte("Something went wrong..."))
// }

