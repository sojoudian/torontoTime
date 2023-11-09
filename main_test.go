package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	//Create a req
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTime)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returend worng status code: got %v want %v", status, http.StatusOK)
	}
	var response TimeResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	_, err = time.Parse("2006-1-2 15:4:5", response.CurrentTime)
	if err != nil {
		t.Errorf("Handler returend unexpected body: got %v", response.CurrentTime)
	}
}
