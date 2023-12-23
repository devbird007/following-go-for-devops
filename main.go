package main

import "fmt"

func main() {
	modelToMake := map[string]string{
		"prius":    "toyota",
		"chevelle": "chevy",
	}

	fmt.Println(modelToMake)

	fmt.Println(modelToMake["vue"])
}
