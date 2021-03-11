package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4)
	fmt.Println(total)

	fmt.Println("-----------")
	writer("Hello World", 1, 2, 3, 4)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func writer(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}
