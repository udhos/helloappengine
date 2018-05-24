package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

const version = "0.2"

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! version %s\n", version)
}
