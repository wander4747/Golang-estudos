package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne, channelTwo := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channelOne <- "Channel one"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channelTwo <- "Channel two"
		}
	}()

	for {
		select {
		case messageOne := <-channelOne:
			fmt.Println(messageOne)
		case messageTwo := <-channelTwo:
			fmt.Println(messageTwo)
		}
	}
}
