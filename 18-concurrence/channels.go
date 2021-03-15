package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go writer("Hello channel", channel)

	//METHOD 1
	/*for {
		message, openChannel := <-channel
		if !openChannel {
			break
		}
		fmt.Println(message)
	}*/

	//METHOD 2
	for message := range channel {
		fmt.Println(message)
	}
}

func writer(text string, channel chan string) {

	for i := 0; i < 5; i++ {
		channel <- fmt.Sprintf("%s %d", text, i)
		time.Sleep(time.Second)
	}

	close(channel)
}
