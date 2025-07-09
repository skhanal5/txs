package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func decode(req interface{}, r *http.Request) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(req)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}
	if err := validate.Struct(req); err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			return fmt.Errorf("validation error: %w", err)
		}
	}
	return err
}

func encode(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return nil
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}
