package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func init() {

	start = time.Now()

}

func service1(c chan string) {
	c <- "Hello from service 1"
}

func service2(c chan string) {
	c <- "Hello from service 2"
}

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1")

	case res := <-chan2:
		fmt.Println("Response from service 2")

	default:
		fmt.Println("No response received")
	}
}
