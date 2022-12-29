/*
	Headers are things that provide additional information about a web request or response.
	They are sent as key-value pairs.
	They are not rendered in the screen, but can affect how things are rendered.
		- Content-Type tells the browser that type of content we're sending back (html, json, etc.)
*/

/*
	https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
	https://pkg.go.dev/net/http#ResponseWriter
	https://pkg.go.dev/net/http#Header
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

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting the server on :5000...")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		panic(err)
	}
}
