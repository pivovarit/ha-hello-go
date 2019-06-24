package main

import "fmt"
import "net/http"
import "time"
import "math/rand"

const GatewayAddress = "gateway:80"

func main() {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		reqType := rand.Intn(3)

		if reqType == 0 {
			url := "http://" + GatewayAddress + "/auth"
			fmt.Println("GET ", url, " at ", time.Now())
			_, _ = http.Get(url)
		} else if reqType == 1 {
			url := "http://" + GatewayAddress + "/users"
			fmt.Println("GET ", url, " at ", time.Now())
			_, _ = http.Get(url)
		} else if reqType == 2 {
			url := "http://" + GatewayAddress + "/"
			fmt.Println("GET ", url, " at ", time.Now())
			_, _ = http.Get(url)
		}
	}
}
