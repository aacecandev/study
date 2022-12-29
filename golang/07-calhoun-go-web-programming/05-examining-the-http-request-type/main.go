/*
	Routing is what happens when you're determining what page to show to the user.
	It varies on different things such as the path, domain, query string, and so on.

	We're going to look at the http.Request type and we're going to see, if we wanted
	to build our own router, what would it look like.
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
	fmt.Fprintf(w, "<h1>You are at %s</h1>", r.URL.Path)
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
