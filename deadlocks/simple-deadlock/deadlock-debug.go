package main

import (
	deadlock "github.com/sasha-s/go-deadlock"
)

var (
	lock1 deadlock.Mutex
	lock2 deadlock.Mutex
)

func process1() {
	for {
		lock1.Lock()
		lock2.Lock()
		lock2.Unlock()
		lock1.Unlock()
	}
}

func process2() {
	for {
		lock2.Lock()
		lock1.Lock()
		lock1.Unlock()
		lock2.Unlock()
	}
}

func main() {
	go process1()
	go process2()
	select {}
}