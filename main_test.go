package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHandler(t *testing.T) {
	config.AWSSvc = "s3"
	config.PresignTime = 5

	req, err := http.NewRequest("GET", "/s3-test-file", nil)
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(httpHandler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusFound, status)
	}
}
