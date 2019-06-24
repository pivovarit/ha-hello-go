package main

import . "fmt"
import "net/http"
import "time"
import "math/rand"

const GatewayAddress = "gateway:80"

func main() {
	for {
		Println("Producing... at ", time.Now())

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://" + GatewayAddress + "/auth")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://" + GatewayAddress + "/users")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		_, _ = http.Get("http://" + GatewayAddress + "/")
	}
}
