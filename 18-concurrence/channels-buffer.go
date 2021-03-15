package main

import "fmt"

func main() {
	channel := make(chan string, 2) //MAXIMUM 2 CHANNEL

	channel <- "Hello World one"
	channel <- "Hello World two"

	message := <-channel
	fmt.Println(message)

	messageTwo := <-channel
	fmt.Println(messageTwo)
}
