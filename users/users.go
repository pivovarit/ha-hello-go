package main

import . "fmt"
import . "net/http"
import "time"

func main() {
	HandleFunc("/api/", func(writer ResponseWriter, req *Request) {
		Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = Fprint(writer, "hello\n")
	})

	HandleFunc("/api/users", func(writer ResponseWriter, req *Request) {
		Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = Fprint(writer, "hello users\n")
	})

	Println("Starting users-go web-server...")
	_ = ListenAndServe(":8080", nil)
}
