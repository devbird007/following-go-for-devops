package main

import "fmt"

func changeValue(word *string) {
	*word += "world"
}

func main() {
	say := "hello"
	changeValue(&say)
	fmt.Println(say)
}
