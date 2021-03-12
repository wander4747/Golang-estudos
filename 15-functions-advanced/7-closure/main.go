package main

import "fmt"

func main() {
	fmt.Println("Closure function")

	text := "Main"

	fmt.Println(text)

	function := closure()
	function()
}

func closure() func() {
	text := "Closure"

	function := func() {
		fmt.Println(text)
	}
	return function
}
