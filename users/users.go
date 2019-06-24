package main

import . "fmt"
import . "net/http"
import "time"

func main() {
	HandleFunc("/api/", func(writer ResponseWriter, req *Request) {
		Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = Fprint(writer, "hello")
	})

	HandleFunc("/api/auth", func(writer ResponseWriter, req *Request) {
		Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = Fprint(writer, "hello auth")
	})

	HandleFunc("/api/users", func(writer ResponseWriter, req *Request) {
		Println("Got a request: ", req.RequestURI, time.Now())

		_, _ = Fprint(writer, "hello users")
	})

	Println("Starting hello-go web-server...")
	_ = ListenAndServe(":8080", nil)
}
