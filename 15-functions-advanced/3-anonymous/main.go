package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()

	func(text string) {
		fmt.Println(text)
	}("Anonymous function with parameter")

	text := func(text string, number int) string {
		return fmt.Sprintf("Receveid -> %s %d", text, number)
	}("Anonymous function with return", 47)

	fmt.Println(text)
}
