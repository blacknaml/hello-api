package rest_test

import (
	"encoding/json"
	"hello-api/handlers/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	handler := http.HandlerFunc(rest.TranslateHandler)

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusOK {
		t.Errorf(`expected status 200 but received %d`, rr.Code)
	}

	var resp rest.Resp
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.Language != "english" {
		t.Errorf(`expected language "english" but received %s`, resp.Language)
	}

	if resp.Translation != "hello" {
		t.Errorf(`expected translation "hello" but received %s`, resp.Translation)
	}

}
