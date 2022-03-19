package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	//Simulates 2 go routines doing some task like hitting an API. Here it is just a ping. But the concept remains the same.
	//Select can be used to wait on a channel to reply , whichever one replies 1st , gets selected.
	//Select can be made nonblocking by adding a default.
	//This also has a timeout defined , if none of the APIs respond within 2 seconds , timeout is reached and we continue.
	go func(c2 chan int) {
		for {
			response, _ := http.Get("https://amazon.com")
			if response.StatusCode == 200 {
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
				c2 <- 1
				break
			}
		}
	}(c2)

	go func(c1 chan int) {
		for {
			response, _ := http.Get("https://google.com")
			if response.StatusCode == 200 {
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
				c1 <- 1
				break
			}
		}
	}(c1)

	select {
	case <-c1:
		fmt.Println("google replied first")
	case <-c2:
		fmt.Println("amazon replied first")
	case <-time.After(time.Second * 2):
		fmt.Println("Timed out.")
	}
	fmt.Println("done")

}
