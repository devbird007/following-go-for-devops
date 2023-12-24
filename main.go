package main

import "fmt"

func main() {
	modelToMake := map[string]string{
		"prius":    "toyota",
		"chevelle": "chevy",
	}

	carMake := modelToMake["chevelle"]
	fmt.Println(carMake)

	if carMake, ok := modelToMake["outback"]; ok {
		fmt.Printf("car model \"outback\" has make %q", carMake)
	} else {
		fmt.Printf("car model \"outback\" has an unknown make\n")
	}

	for key, val := range modelToMake {
		fmt.Printf("car model %q has make %q\n", key, val)
	}

	var x int = 23
	fmt.Println(&x)
}
