package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloworldHandler)
	http.ListenAndServe(":8080", nil)
}

//HelloworldHandler HelloWorld net/http handler
func HelloworldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}
