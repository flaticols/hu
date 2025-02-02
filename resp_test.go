package hu

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// A type that is not JSON serializable (channels cannot be marshalled).
type nonJSONable struct {
	Ch chan int
}

func TestRespJSON_Success(t *testing.T) {
	// Arrange: create a ResponseRecorder and a simple payload.
	rr := httptest.NewRecorder()
	payload := map[string]string{"message": "hello world"}

	// Act: call RespJSON with a test code.
	err := RespJSON(rr, http.StatusOK, payload)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Assert: Content-Type header, status code, and payload.
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", ct)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{"message":"hello world"}`
	if strings.TrimSpace(rr.Body.String()) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, rr.Body.String())
	}
}

func TestRespJSON_ErrorMarshalling(t *testing.T) {
	// Arrange: Create a response recorder with data that cannot be marshalled to JSON.
	rr := httptest.NewRecorder()
	payload := nonJSONable{Ch: make(chan int)}

	// Act: calling RespJSON should return an error.
	err := RespJSON(rr, http.StatusOK, payload)
	if err == nil {
		t.Fatal("Expected an error when marshalling nonJSONable data, got nil")
	}

	if !errors.Is(err, err) { // Just a dummy check to use the error; you can check error string if desired.
		t.Errorf("Expected marshalling error, got %v", err)
	}
}

func TestResp(t *testing.T) {
	// Arrange: Create a response recorder.
	rr := httptest.NewRecorder()

	// Act: call Resp.
	Resp(rr, http.StatusNoContent)

	// Assert: status code set; no body modifications.
	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rr.Code)
	}
}

func TestRespBad(t *testing.T) {
	// Arrange: create a ResponseRecorder and payload.
	rr := httptest.NewRecorder()
	payload := map[string]string{"error": "bad request"}

	// Act: calling RespBad should emit a bad request.
	err := RespBad(rr, payload)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Assert: Check status code and content type.
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}

	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", ct)
	}

	expectedBody := `{"error":"bad request"}`
	if strings.TrimSpace(rr.Body.String()) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, rr.Body.String())
	}
}
