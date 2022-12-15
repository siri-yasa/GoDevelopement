package main

import (
	"errors"
	"fmt"
)

type err struct {
}

func (e err) Error() string {
	return "error happened"
}
func main() {
	err1 := err{}
	err2 := e2()
	if err1 == err2 {
		fmt.Println("Equality operator:Both errors are equal")
	} else {
		fmt.Println("Equality operator:Both errors are not equal")
	}
	if errors.Is(err2, err1) {
		fmt.Println("Is function: Both errors are equal")
	}
}
func e2() error {
	return fmt.Errorf("e2:", err{})
}

// Output
// Equality Operator: Both errors are not equal
// Is function: Both errors are equal
