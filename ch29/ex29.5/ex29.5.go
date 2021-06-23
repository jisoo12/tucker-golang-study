package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World")
	})
	mux.HandleFunc("/bar", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello Bar")
	})
	
	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}