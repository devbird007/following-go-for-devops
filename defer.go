package main

import (
	"fmt"
)

func printStuff() (value string) {
	// This will happen last
	defer fmt.Println("exiting")

	// This anonymous function sets our named return value. This
	// can only be done with named returns.
	defer func() {
		value = "we returned this"
	}()

	// This will happen first.
	fmt.Println("I am printing stuff")

	// We return an empty string, but that defer'd anonymous function
	// is going to intercept this and change the value.
	return ""
}

func main() {
	v := printStuff()
	fmt.Println(v)
}
