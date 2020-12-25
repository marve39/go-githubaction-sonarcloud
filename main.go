package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//main main method
func main() {
	startHTTPServer(false)
}

func startHTTPServer(isTest bool) {
	if !isTest {
		http.HandleFunc("/", HelloworldHandler)
		http.HandleFunc("/version", VersionHandler)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}

//HelloworldHandler HelloWorld net/http handler
func HelloworldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

//VersionHandler method to get version of application
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("version.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", content)
}
