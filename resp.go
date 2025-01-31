package hu

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respJSON(w http.ResponseWriter, code int, data interface{}) error {
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

func resp(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func respBad(w http.ResponseWriter, data interface{}) error {
	return respJSON(w, http.StatusBadRequest, data)
}
