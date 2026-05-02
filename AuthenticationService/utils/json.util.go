package utils

import (
	"AuthenticationService/model"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate



func init() {
	Validator = newValidator()
}

func newValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}


func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, model.ErrorResponse{Message: message})
}

func ReadJSON(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}
