package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	We are naming the func handlerFunc because net/http has a type HandlerFunc
*/

func handlerFunc(request http.ResponseWriter, response *http.Request) {
	/*
		The ResponseWriter is an interface that defines a set of methods that allow us to write data to the response. It can write headers, set cookies, set http status codes...
	*/

	/*
		The Request is a pointer to and explicit struct type, is not an interface. It contains all the information about the request that was made to the server and that we can read (headers, cookies, body, etc.)
	*/

	/*
		Why one is an interface and the other a struct?

		Since ResponseWriter is an interface, it can be substituted with multiple implementations (JSON, gRPC, test mocks, etc.).
		Different types of connections might allow for different types of responses. For example, a websocket connection might allow for a different type of response than a regular HTTP connection.
		The net/http/httptest package has the ResponseRecorder type that is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.


		The Request is a struct because it is a specific type of data that we want to work with.
	*/
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	// fmt.Fprint(os.Stdout, "Request received")
}

func main() {
	// First part, where we define handlers
	http.HandleFunc("/", handlerFunc)
	time.Sleep(3 * time.Second)

	// Second part, where we start the server
	fmt.Println("Starting the server on :5000...")
	err := http.ListenAndServe(":5000", nil) // nil is the default ServeMux handler

	if err != nil {
		panic(err)
	}
}

// This is the same

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", handlerFunc)
// 	fmt.Println("Starting the server on :5000...")
// 	http.ListenAndServe(":5000", nil)
// }
