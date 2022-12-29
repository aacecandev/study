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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
	<li>
		<b>Question 1</b>
		Answer 1
	</li>
	<li>
		<b>Question 2</b>
		Answer 2
	</li>
	<li>
		<b>Question 3</b>
		Answer 3
	</li>
</ul>
`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

type Router struct{}

//	Using this approach we don't worry about:
//	 - HTTP methods
//		 - GET /galleries
//		 - POST /galleries
//	 - dynamic parameters in the path
//		 - GET /galleries/:id
//		 - GET /galleries/:id/edit
//	 - prefix matching
//
//	And using the standard library, specifically the ServeMux method, makes this
// 	way more harder than using 3rd party libraries like Gin or Gorilla
func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :5000...")
	err := http.ListenAndServe(":5000", router)

	if err != nil {
		panic(err)
	}
}
