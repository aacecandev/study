package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string        // this encodes the script
	Bio  template.HTML // this leaves the script as it is
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Bio:  `<script>alert("XSS")</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
