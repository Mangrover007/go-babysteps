package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"

	"fmt"
	"os"
	"log/slog"
	"time"
	"errors"
)

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	var StatusCode int

	defer func() {
		logger.Info("handled request", "StatusCode", StatusCode, "method", r.Method, "path", "/multiply", "duration", time.Since(start))
	}()

	if r.Method != http.MethodPost {
		StatusCode = http.StatusBadRequest
		response.BadRequest[int](w, "Method must be a POST request")
		return
	}

	var body models.RequestBodyOther
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		var typeError *json.UnmarshalTypeError

		if errors.As(err, &typeError) {
			fmt.Printf("%w\n", err)
			StatusCode = http.StatusBadRequest
			response.BadRequest[int](w, "cannot multiply strings send NUMBERS")
			return
		}

		StatusCode = http.StatusInternalServerError
		return
	}

	StatusCode = http.StatusOK 
	response.Success[int64](w, int64(body.Num1) * int64(body.Num2))
}

