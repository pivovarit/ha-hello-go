package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = fmt.Fprint(writer, "hello")
	})

	http.HandleFunc("/api/auth", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = fmt.Fprint(writer, "hello auth")
	})

	http.HandleFunc("/api/users", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = fmt.Fprint(writer, "hello users")
	})

	fmt.Println("Starting hello-go web-server...")
	_ = http.ListenAndServe(":8080", nil)
}
