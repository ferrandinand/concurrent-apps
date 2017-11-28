package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func main2() {
	wait.Add(1)

	toUpperASync("Hello Callbacks!", func(v string) {
		fmt.Printf("Callback: %s\n", v)
		wait.Done()
	})

	fmt.Println("Waiting async response")
	wait.Wait()
}
func toUpperASync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}
