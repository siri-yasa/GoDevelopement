package main

import (
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}
func mul(a int, b int) int {
	return a * b
}
func main() {
	sumResult := sum(12, 23)
	mulResult := mul(14, 13)
	fmt.Println(sumResult)
	fmt.Println(mulResult)
}
