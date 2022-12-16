package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)

	}
}
func main() {
	go hello("best")
	hello("super")
}
