// Sample program to show how a bytes.Buffer can also be used
// with the io.Copy function.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main is the entry point for the application.
func main() {
	var b bytes.Buffer

	// Write a string to the buffer.
	b.Write([]byte("Hello"))

	// Use Fprintf to concatenate a string to the Buffer.
	//  accepts an interface value of type io.Writer as its first parameter. Since pointers of type bytes.Buffer implement the io.Writer interface

	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to stdout.
	// Since pointers of type bytes.Buffer also implement the io.Reader interface, the io.Copy function can be used to display the contents of the buffer to the terminal window.
	io.Copy(os.Stdout, &b)
}
