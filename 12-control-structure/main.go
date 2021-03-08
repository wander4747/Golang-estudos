package main

import "fmt"

func main() {
	fmt.Println("Control structure")

	number := 10

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less than or equal to 15")
	}

	if number2 := 10; number2 > 0 {
		fmt.Println("Number greater than zero")
	}
}
