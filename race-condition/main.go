package main

import (
	"fmt"
	"sync"
)

// Teste based on this article
// https://golangbot.com/mutex

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}

	w.Wait()
	fmt.Println("final value of x", x)
}
