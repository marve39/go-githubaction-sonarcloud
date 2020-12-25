package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testingWord = "hello"

//TestHelloworldHandler Test Handler for hello world
func TestHelloworldHandler(t *testing.T) {
	r := httptest.NewRequest("GET", fmt.Sprintf("/%s", testingWord), nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(HelloworldHandler)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	if s := w.Code; s != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v, want %v", s, http.StatusOK)
	}

	// check HTTP response body
	want := testingWord
	if w.Body.String() != want {
		t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), want)
	}
}
