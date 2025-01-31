package hu

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespJSON(w http.ResponseWriter, code int, data interface{}) error {
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

func RespBad(w http.ResponseWriter, data interface{}) error {
	return RespJSON(w, http.StatusBadRequest, data)
}
