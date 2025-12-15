package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"

	"log/slog"
	"time"
	"os"
	"errors"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	var StatusCode int

	defer func() {
		logger.Info("handled request", "StatusCode", StatusCode, "method", r.Method, "path", "/divide", "duration", time.Since(start))
	}()

	if r.Method != http.MethodPost {
		StatusCode = http.StatusBadRequest
		response.BadRequest[int](w, "Method must be a POST request")
		return
	}

	var body models.RequestBodyDivide
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		var typeError *json.UnmarshalTypeError

		if errors.As(err, &typeError) {
			StatusCode = http.StatusBadRequest
			response.BadRequest[int](w, "cannot divide strings send NUMBERS")
			return
		}

		StatusCode = http.StatusInternalServerError
		return
	}

	if body.Divisor == float64(0) {
		StatusCode = http.StatusBadRequest
		response.BadRequest[int](w, "Cannot divide by 0")
		return
	}

	StatusCode = http.StatusOK
	response.Success[float64](w, body.Dividend/body.Divisor)
}
