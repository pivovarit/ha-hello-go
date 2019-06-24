package main

import . "fmt"
import "net/http"
import "time"
import "math/rand"

func main() {
	for {
		Println("Producing... at ", time.Now())

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://nginx:80/api/auth")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://nginx:80/api/users")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://nginx:80/")
	}
}
