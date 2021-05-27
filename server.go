package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// Handler for taking "/" uses http.DefaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted) // Returns 202 status
		parameters := r.URL.Path[1:]
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(parameters))
		// working with parameters & differnt request types
		if r.Method == "GET" {
			for k, v := range r.URL.Query() {
				// print parameters to log to see
				log.Println(string(k) + " :" + v[0])
				// want to use the schema package to easily turn into a struct
			}
		}

	})
	// Basic Handler to see what happens
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bye!!")
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil)) // passing nil to ListenAndServe() does this manually
}
