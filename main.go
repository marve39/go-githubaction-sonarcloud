package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//main main method
func main() {
	startHTTPServer(false)
}

func startHTTPServer(isTest bool) {
	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", HelloworldHandler)
	http.HandleFunc("/version", VersionHandler)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- srv.ListenAndServe()
	}()

	if isTest {
		log.Println("Test")
		srv.Shutdown(context.Background())

		//errs <- fmt.Errorf("%s", "Test Only")
	}

	log.Print("ListenAndServe: ", <-errs)
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
