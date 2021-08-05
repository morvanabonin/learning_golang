package main

import (
	"fmt"
	"time"
)

// Teste based on this article
// https://golangbot.com/goroutines

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Finish him!")

}
