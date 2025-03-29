package hu

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RespJSON sends a JSON response with the specified status code and data.
// It sets the Content-Type header to application/json and marshals the data to JSON.
// If marshaling or writing the response fails, it returns an error.
func RespJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}

	w.WriteHeader(code)
	if _, err = w.Write(payload); err != nil {
		return fmt.Errorf("error writing response: %w", err)
	}

	return nil
}

func Resp(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func RespBad(w http.ResponseWriter, data any) error {
	return RespJSON(w, http.StatusBadRequest, data)
}
