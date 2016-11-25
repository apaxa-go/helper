package httph

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthError(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		AuthError(w, nil)
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("TestAuthError: should set status code to unauthorized, but got %v", w.Code)
	}
	if strings.TrimRight(string(w.Body.Bytes()), " \n") != message {
		t.Errorf("TestAuthError: should return '%v', but got '%v'", message, w.Body.String())
	}
}
