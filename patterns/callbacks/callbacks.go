package main

import (
	"fmt"
	"strings"
)

/*
func toUpperSync(word string) string {}

func toUpperSync(word string, f func(string)) {}

func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}
*/

func main2() {
	toUpperSync("Hello Callbacks!", func(v string) {
		fmt.Printf("Callback: %s\n", v)
	})
}

func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}
