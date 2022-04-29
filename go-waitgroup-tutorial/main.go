package main

/*
 * WaitGroup tutorial
 *
 * References by:
 * https://tutorialedge.net/golang/go-waitgroup-tutorial
 */
import (
	"fmt"
	"sync"
)

// func myFunc() {
	// fmt.Println("Inside my goroutine")
// }

// func main() {
	// fmt.Println("Hello World")
	// go myFunc()
	// fmt.Println("Finished Execution")
// }

func main() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go func() {
		fmt.Println("Inside my goroutine")
		waitgroup.Done()
	}()
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}