package main

import "fmt"

func main() {
	fmt.Println("Recursive function")
	position := uint(20)
	fmt.Println(fibonacci(position))
}

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}
