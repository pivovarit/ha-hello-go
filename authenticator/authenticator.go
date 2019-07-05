package main

import . "fmt"
import . "net/http"
import "time"

func main() {
	HandleFunc("/api/auth", func(writer ResponseWriter, req *Request) {
		Println("Authenticating: ", req.RequestURI, time.Now())

		_, _ = Fprint(writer, "hello auth")
	})

	Println("Starting authenticator-go web-server...")
	_ = ListenAndServe(":8080", nil)
}
