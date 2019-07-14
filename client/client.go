package main

import (
	"fmt"
	"strconv"
)
import "net/http"
import "time"
import "math/rand"

const GatewayAddress = "gateway:80"

func main() {
	for {
		time.Sleep(time.Duration(1000+rand.Intn(1000)) * time.Millisecond)

		fmt.Println("")
		fmt.Println("### PRODUCING ###")
		fmt.Println("")

		reqType := rand.Intn(3)

		if reqType == 0 {
			url := "http://" + GatewayAddress + "/auth"
			fmt.Println("GET ", url, " at ", time.Now())
			_, _ = http.Get(url)
		} else if reqType == 1 {
			url := "http://" + GatewayAddress + "/users" + "/" + strconv.FormatInt(int64(rand.Intn(100000)), 10)
			fmt.Println("GET ", url, " at ", time.Now())
			_, _ = http.Get(url)
		} else if reqType == 2 {
			url := "http://" + GatewayAddress + "/payments"
			fmt.Println("GET ", url, " at ", time.Now())
			_, _ = http.Get(url)

		}
	}
}
