/*
	We're going to proceed to build the router using the r.URL.Path value.
*/

/*
	https://pkg.go.dev/net/http#Request
	https://pkg.go.dev/net/url#URL
*/

package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// The difference between w.Header().Add() and w.Header().Set() is that Add() will add a new value to the header, while Set() will overwrite the existing value.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch, email me at <a href=\"mailto:dev@aacecan.com\">dev@aacecan.com</a>.</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO: handle the page not found error
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting the server on :5000...")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		panic(err)
	}
}
