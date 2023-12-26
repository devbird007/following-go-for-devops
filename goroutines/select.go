package main

import (
	"fmt"
)

func main() {
	for {
		select {
		case v := <-inCh1:
			go fmt.Println("received(inCh1): ", v)
		case v := <-inCh2:
			go fmt.Println("received(inCh2): ", v)
		}

		// This is for an alternative select statement where you
		// want to check if a value is on a channel, but if it is
		// not, then you want to move on to another thing by
		// executing a default statement. Unlike the previous
		// select statement that will wait for channel data
		// indefinitely.

		select {
		case s := <-ch:
			fmt.Printf("had a string(%s) on the channel\n", s)
		default:
			fmt.Println("channel was empty")
		}
	}
}
