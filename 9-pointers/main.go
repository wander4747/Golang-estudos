package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var number int = 10
	var numberAxu = number

	fmt.Println(number, numberAxu)

	number++
	fmt.Println(number, numberAxu)

	//POINTER
	var number2 int
	var pointer *int

	number2 = 100
	pointer = &number2

	fmt.Println(number2, pointer)
	fmt.Println(number2, *pointer)
}
