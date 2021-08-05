package main

import (
	"fmt"
	"time"
)

// Test based on this article
// https://golangbot.com/mutex

func plus() {
	var x int
	x = x + 1
	fmt.Println("o valor de x => ", x)
}

func plus2() {
	var x int
	x = x + 1
	fmt.Println("o valor de x => ", x)
}

func plus3() {
	var x int
	x = x + 1
	fmt.Println("o valor de x => ", x)
}

func main() {
	go plus()
	go plus2()
	go plus3()
	time.Sleep(1 * time.Second)
	fmt.Println("Finish him!")
}
