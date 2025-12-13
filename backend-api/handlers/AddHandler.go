package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/Mangrover007/go-babysteps/backend-api/models"
	"github.com/Mangrover007/go-babysteps/backend-api/response"
)

type AddHandler struct {}

func (_ *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// any models.ResultType would be ok here I think
		response.BadRequest[int](w, "Method must be a POST request")
		return
	}
	
	var body models.RequestBodyOther
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.BadRequest[int](w, err.Error())
		return
	}
	
	response.Success[int](w, body.Num1 + body.Num2)
}

