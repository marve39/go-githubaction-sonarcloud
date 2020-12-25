package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testingWord = "hello"

//TestStartHTTPServer Test HTTPServer
func TestStartHTTPServer(t *testing.T) {
	startHTTPServer(true)
}

//TestHelloworldHandler Test Handler for hello world
func TestHelloworldHandler(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	httpTest(t, testingWord, testingWord, HelloworldHandler)
}

//TestMain Test main function
func TestMain(t *testing.T) {
	go main()
	log.Printf("Tested")
}

//TestVersionHandler Test version handler
func TestVersionHandler(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	httpTest(t, "version", "dummy", VersionHandler)
}

//httpTest method for doing http test
func httpTest(t *testing.T, url string, word string, handler func(http.ResponseWriter, *http.Request)) {
	r := httptest.NewRequest("GET", fmt.Sprintf("/%s", url), nil)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(handler)

	h.ServeHTTP(w, r)

	// check HTTP response status code
	if s := w.Code; s != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v, want %v", s, http.StatusOK)
	}

	// check HTTP response body
	want := word
	if w.Body.String() != want {
		t.Errorf("handler return wrong status code: got %v, want %v", w.Body.String(), want)
	}
}
