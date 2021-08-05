package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	now := time.Now()
	fmt.Println(now)

	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	fmt.Println(formatted)
}
