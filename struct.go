package main

import (
	"fmt"
)

type Record struct {
	Name string
	Age  int
}

// Pointer-value function
func changeName(r *Record) {
	r.Name = "Peter"
	fmt.Println("Inside changeName: ", r.Name)
}

// Constructor
func NewRecord(name string, age int) (*Record, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be the empty string")
	}

	if age <= 0 {
		return nil, fmt.Errorf("age cannot be <=0")
	}
	return &Record{Name: name, Age: age}, nil
}

func main() {

	// changing a record name
	rec := &Record{Name: "John"}
	changeName(rec)
	fmt.Println("Main: ", rec.Name)

	rus, err := NewRecord("Tim", 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rus)
}
