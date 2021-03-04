package main

import "fmt"

func main() {
	sum := sum(5, 2)
	fmt.Println(sum)

	var f = func(text string) string {
		return text
	}

	result := f("Function type")
	fmt.Println(result)

	resultSum, resultSubtraction := calculate(10, 15)
	fmt.Println(resultSum, resultSubtraction)
}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

//more than one return
func calculate(n1, n2 int) (int, int) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}
