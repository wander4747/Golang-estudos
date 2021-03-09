package main

import "fmt"

func main() {
	sum, subtraction := calculate(15, 10)
	fmt.Println(sum)
	fmt.Println(subtraction)
}

func calculate(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}
