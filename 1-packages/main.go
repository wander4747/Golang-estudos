package main

import (
	"fmt"
	"module/auxiliar"
)

func main() {
	fmt.Println("Run: go run main.go")
	fmt.Println("--------------")
	fmt.Println("Run: go build")
	fmt.Println("generate exe")
	fmt.Println("--------------")
	fmt.Println("./module : run exe")

	auxiliar.Write()
}
