package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := 10
	for {
		n := rand.Intn(10) + 1
		if n < number {
			fmt.Printf("%v слишком маленькое число.\n", n)
		} else if n > number {
			fmt.Printf("%v слишком большое число.\n", n)
		} else {
			fmt.Printf("Угадал! %v\n", n)
			break
		}
	}
}
