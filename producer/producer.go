package main

import "net/http"
import "time"
import "math/rand"

func main() {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://haproxy:8080/api/auth")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://haproxy:8080/api/users")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://haproxy:8080/")
	}
}
