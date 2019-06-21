package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println("Got a request: " + req.RequestURI)

		_, _ = fmt.Fprint(writer, "hello world")
	})

	fmt.Println("Starting hello-go web-server...")
	_ = http.ListenAndServe(":8080", nil)
}
