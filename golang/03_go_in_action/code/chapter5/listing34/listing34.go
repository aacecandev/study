// Sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init is called before main.
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

// main is the entry point for the application.
func main() {
	// Get a response from the web server.
	// Return a pointer of type http.Request
	// http.Request contains a field called Body which is an interface of
	// type io.ReadCloser
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Copies from the Body to Stdout.
	// accepts values of interface type io.Reader for its second parameter, and this value represents a source of data to stream from
	// The first parameter for io.Copy represents the destination and must be a value that implements the io.Writer interface
	// When we pass the Body and Stdout values to the io.Copy function, the function streams data from the web server to the terminal window in small chunks
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
