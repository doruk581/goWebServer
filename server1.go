package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	var host = "localhost:"
	var port = "8000"

	http.HandleFunc("/", handler) // each request call handler
	log.Fatal(http.ListenAndServe(host+port, nil))

}

//handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
