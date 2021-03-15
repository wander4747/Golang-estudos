package main

import (
	"fmt"
	"time"
)

func main() {
	go writer("Hello world!")
	writer("Programming in Go!")
}

func writer(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
