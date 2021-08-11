package main

import (
	"fmt"
	"time"
)

// Test based on this article
// https://golangbot.com/goroutines

func hello() {
	fmt.Println("Hello world goroutine")
}

func helloAgain() {
	fmt.Println("Hello Again!")
}

func helloOneMoreTime() {
	fmt.Println("Hello one more time!")
}

// Goroutines execurte without an ordered
func main() {
	go hello()
	go helloAgain()
	go helloOneMoreTime()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
