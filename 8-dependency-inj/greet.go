package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// How to test? -> Printf prints to stdout, hard to capture using the testing framework
// We need to inject. (fancy word for pass in) the dependency of printing
// func Greet(name string) {
// 	fmt.Printf("Hello, %s", name)
// }

// Use the writer to send the greeting to the buffer in our test.
// Fprintf is like Printf but insteads takes a Writer to send the string to,
// whereas Printf defaults to stdout
// v2
// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// Fprintf allows you to pass an io.Writer which we know os.Stdout and bytes.Buffer
// both implement. If we change our code to use the more general purpose interface
// we can now use it in both tests and our application
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//More on io.Writer

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// Why does this work?
// When you write an HTTP handler, you are giveen an http.ResponseWriter and the
// http.Request that was used to mkae the requestt. When you implement your
// server you WRITE your response using the writer

// http.ResponseWriter implements io.Writer
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
