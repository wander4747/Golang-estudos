package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(writer("Hello world"), writer("Programming in GO"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channelOne, channelTwo <-chan string) <-chan string {
	channelOut := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelOne:
				channelOut <- message
			case message := <-channelTwo:
				channelOut <- message
			}
		}
	}()

	return channelOut
}

func writer(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
