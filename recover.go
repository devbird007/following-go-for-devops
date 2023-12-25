package main

import (
	"log"
)

func someFunc() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("called recover, panic was: %q", r)
		}
	}()

	panic("oh no!!!")
}

func main() {
	someFunc()
}
