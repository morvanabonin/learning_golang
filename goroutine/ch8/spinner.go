package main

import time

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) //lento
	fmt.Printf("\rFibonacci(%d) = %d\n", n. fibN)
}

func spinner(delay time.Duration) {
	for {
		// this string in range is used to go through the slice
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}