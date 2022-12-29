package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my great site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)

	fmt.Println("Starting the server on :5000...")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		panic(err)
	}
}
