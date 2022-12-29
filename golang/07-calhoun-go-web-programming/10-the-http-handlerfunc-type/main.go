package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch, email me at <a href=\"mailto:dev@aacecan.com\">dev@aacecan.com</a>.</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

func main() {
	// var router http.HandlerFunc = pathHandler
	fmt.Println("Starting the server on :5000...")
	// http.HandleFunc(pathHandler) is a type conversion, like int64(42), not a function call
	err := http.ListenAndServe(":5000", http.HandlerFunc(pathHandler))

	// http.Handler - interface with the ServeHTTP method
	// http.HandlerFunc - function type that implements the ServeHTTP method. Accepts same args as ServeHTTP method. Also implements http.Handler interface

	// http.Handle("/", router) - http.Handle accepts a path and a handler. The handler must implement the http.Handler interface
	// http.HandleFunc("/", pathHandler) - http.HandleFunc accepts a path and a handler function. The handler function must accept a ResponseWriter and a Request as args and return nothing

	if err != nil {
		panic(err)
	}
}
