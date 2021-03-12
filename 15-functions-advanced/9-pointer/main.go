package main

import "fmt"

func main() {
	number := 20
	inverterNumber := inverterSignal(number)
	fmt.Println(inverterNumber)

	number2 := 40
	inverterSignalWithPointer(&number2)
	fmt.Println(number2)
}

func inverterSignal(number int) int {
	return number * -1
}

func inverterSignalWithPointer(number *int) {
	*number = *number * -1
}
