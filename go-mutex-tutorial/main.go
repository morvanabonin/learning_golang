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
	"reflect"
)

var (
	mutex sync.Mutex
	balance int
)

const mutexLocked = 1

func deposit(value int, wg *sync.WaitGroup) {
	fmt.Println("mutex locked =", MutexLocked(&mutex))
	mutex.Lock()
	fmt.Println("mutex locked =", MutexLocked(&mutex))
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	fmt.Println("mutex locked =", MutexLocked(&mutex))
	wg.Done()

}

func withdraw(value int, wg *sync.WaitGroup) {
	fmt.Println("mutex locked =", MutexLocked(&mutex))
	mutex.Lock()
	fmt.Println("mutex locked =", MutexLocked(&mutex))
	fmt.Printf("Withdrawing %d from account with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	fmt.Println("mutex locked =", MutexLocked(&mutex))
	wg.Done()
}

func MutexLocked(m *sync.Mutex) bool {
    state := reflect.ValueOf(m).Elem().FieldByName("state")
    return state.Int()&mutexLocked == mutexLocked
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





