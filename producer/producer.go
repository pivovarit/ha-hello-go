package main

import (
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(100 * time.Millisecond)
		_, _ = http.Get("http://haproxy:8080/")
	}
}
