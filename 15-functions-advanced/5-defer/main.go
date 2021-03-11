package main

import "fmt"

func main() {
	defer textOne()
	textTwo()
	textThree()
}

func textOne() {
	fmt.Println("Run textOne")
}

func textTwo() {
	fmt.Println("Run textTwo")
}

func textThree() {
	defer fmt.Println("Run textThree defer")
	fmt.Println("Run textThree")
}
