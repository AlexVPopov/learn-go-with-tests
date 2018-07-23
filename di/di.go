package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet writes a greeting to the writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHanlder handles greetings
func MyGreeterHanlder(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHanlder))
}
