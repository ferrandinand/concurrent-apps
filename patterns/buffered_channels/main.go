package main

import (
	"fmt"
	"time"
)

func main() {
	//With channel capacity
	channel := make(chan string, 1)

	go func(ch chan<- string) {
		ch <- "Hellow world 1"
		fmt.Println("Finishing goroutine")
	}(channel)
	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
}
