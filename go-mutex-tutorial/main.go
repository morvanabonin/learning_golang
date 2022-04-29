package main

/*
 * Mutex tutorial
 *
 * References by:
 * https://tutorialedge.net/golang/go-mutex-tutorial
 */
import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()

}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Hello World")

	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}





