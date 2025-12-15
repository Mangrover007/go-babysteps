package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"

	"log/slog"
	"time"
	"os"
	"fmt"
	"errors"
)

func SumHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	var StatusCode int

	defer func() {
		logger.Info("handled request", "StatusCode", StatusCode, "method", r.Method, "path", "/sum", "duration", time.Since(start))
	}()

	if r.Method != http.MethodPost {
		StatusCode = http.StatusBadRequest
		response.BadRequest[int](w, "Method must be a POST request")
		return
	}

	var body models.RequestBodySum
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		var typeError *json.UnmarshalTypeError

		if errors.As(err, &typeError) {
			fmt.Printf("%w\n", err)
			StatusCode = http.StatusBadRequest
			response.BadRequest[int](w, "cannot add strings send NUMBERS")
			return
		}

		StatusCode = http.StatusInternalServerError
		return
	}

	if body.Items == nil {
		StatusCode = http.StatusBadRequest
		response.BadRequest[int](w, "\"items\" not found in request body")
		return
	}

	var sum int64 = 0
	for _, num := range body.Items {
		sum += int64(num)
	}

	StatusCode = http.StatusOK
	response.Success[int64](w, sum)
}

