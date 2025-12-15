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

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	var StatusCode int

	defer func() {
		logger.Info("handled request", "StatusCode", StatusCode, "method", r.Method, "path", "/subtract", "duration", time.Since(start))
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
			StatusCode = http.StatusBadRequest
			response.BadRequest[int](w, "cannot subtract strings send NUMBERS")
			return
		}
		
		StatusCode = http.StatusInternalServerError
		return
	}

	StatusCode = http.StatusOK
	response.Success[int](w, body.Num1 - body.Num2)
}

