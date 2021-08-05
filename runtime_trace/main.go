package main

import (
	"os"
	"runtime/trace"
)

// Test based on this article
// https://blog.gopheracademy.com/advent-2017/go-execution-tracer/

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	// create new channel of type int
	ch := make(chan int)

	// start new anonymos goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()

	// read from channel
	<-ch

}
