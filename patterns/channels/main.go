package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		channel <- "Hellow world"
		fmt.Println("Finishing goroutine")
		waitGroup.Done()
	}()
	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}
