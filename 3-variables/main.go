package main

import "fmt"

func main() {
	//MODEL 1
	var name string = "Wander"
	lastname := "Douglas" // explicit

	fmt.Println(name, lastname)

	//MODEL 2
	var (
		car string = "Ferrari"
		bus string = "Mercedez"
	)

	fmt.Println(car, bus)

	//MODEL 3
	brother, sister := "Joe", "Mary"
	fmt.Println(brother, sister)

}
