package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait2 sync.WaitGroup

func main() {
	wait2.Add(1)

	toUpperASyncHell("First Callbacks!", func(v string) {
		toUpperASyncHell("Second Callbacks!", func(v string) {

			fmt.Printf("Callback: %s\n", v)
			wait2.Done()
		})
	})

	fmt.Println("Waiting async response")
	wait2.Wait()
}
func toUpperASyncHell(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}
