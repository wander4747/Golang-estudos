package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		writer("Hello world!")
		waitGroup.Done()
	}()

	go func() {
		writer("Programming in Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func writer(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
