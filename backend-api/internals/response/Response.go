package response

import (
	"net/http"
	"encoding/json"

	"github.com/Mangrover007/go-babysteps/backend-api/models"
)

func NotAuthorized(w http.ResponseWriter, message string) error {
	w.WriteHeader(http.StatusUnauthorized)

	res := models.Response[string]{
		Data: "",
		Error: message,
	}

	err := json.NewEncoder(w).Encode(res)
	return err
}

func Success[T models.ResponseData] (w http.ResponseWriter, data T) error {
	w. WriteHeader(http.StatusOK)

	res := models.Response[T]{
		Data: data,
		Error: "",
	}
	
	err := json.NewEncoder(w).Encode(res)
	return err
}

func BadRequest(w http.ResponseWriter, message string) error {
	w.WriteHeader(http.StatusBadRequest)

	res := models.Response[string]{
		Data: "",
		Error: message,
	}

	err := json.NewEncoder(w).Encode(res)
	return err
}

func InternalServerError(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusInternalServerError)

	res := models.Response[string]{
		Data: "",
		Error: "Something went wrong",
	}
	
	err := json.NewEncoder(w).Encode(res)
	return err
}

func NotFound(w http.ResponseWriter, message string) error {
	w.WriteHeader(http.StatusNotFound)

	res := models.Response[string]{
		Data: "",
		Error: message,
	}

	err := json.NewEncoder(w).Encode(res)
	return err
}

