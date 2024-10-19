package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler
func baseRoute(w http.ResponseWriter, r *http.Request) {
	/**
	because '/' is a fixed path it can match with all urls
	to prevent this behaviour you can check the path of request
	if != '/' throw 404
	*/

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to snippet"))
}
func viewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a single snippet"))
}
func createSnippets(w http.ResponseWriter, r *http.Request) {
	/*
		to allow only post requests since servemux doesnt support restricting
		request methods to endpoints
	*/

	if r.Method != "POST" {

		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create a snippet"))
}
func main() {
	fmt.Println("Hello world")

	// using server mux to rout
	mux := http.NewServeMux()
	mux.HandleFunc("/", baseRoute)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippets)

	// server
	log.Print("Starting server on :9090")
	err := http.ListenAndServe(":9090", mux)
	log.Fatal(err)
}
