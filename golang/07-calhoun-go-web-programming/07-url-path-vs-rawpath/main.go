/*
	In URLs there are some special characters
		- / (forward slash) is used to separate the path segments
		- ? (question mark) -> provide key=value pairs
		- & (ampersand) -> separate key=value pairs
		- # (hash) -> go to a section of the HTML code (and id attribute)
		- % (percent sign) -> escape characters
		- = (equals sign) -> separate key from value
		- + (plus sign) -> space character
		- : (colon) -> separate the port number from the host name
		- ; (semicolon) -> separate the parameters from the URL
		- @ (at sign) -> separate the user name from the host name

	r.URL.Path -> has the decoded path, where ?key=value is not part of the path
	r.URL.RawPath -> has the encoded path, where ?key=value is part of the path
*/

/*
	https://meyerweb.com/eric/tools/dencoder/
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

	// Path string // path (relative paths may omit leading slash)
	// RawPath string // encoded path hint (see EscapedPath method)
	// I

	// http://localhost:PORT/this%20is%20a%20test

	// localhost:5000/dog%2Fcat%3Fkey=value
	// /dog/cat?key=value
	// /dog%2Fcat%3Fkey=value

	// with RawPath we can see if the path was actually encoded or not
	// fmt.Fprintln(w, r.URL.Path)
	// fmt.Fprintln(w, r.URL.RawPath)

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
