// Tests based on this article
// https://zetcode.com/golang/byte/

package main

import "fmt"

func main() {
	var x1 byte = 84
	var x2 byte = 105
	var x3 byte = 109
	var x4 byte = 101
	var x5 byte = 210
	var x6 byte = 93

	s1 := string([]byte{65, 99, 116, 105, 111, 110})
	s2 := string([]byte{67, 111, 110, 116, 97, 99, 116})
	s3 := string([]byte{84, 105, 109, 101})

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)
	fmt.Println(x6)

	fmt.Printf("%c\n", x1)
	fmt.Printf("%c\n", x2)
	fmt.Printf("%c\n", x3)
	fmt.Printf("%c\n", x4)
	fmt.Printf("%c\n", x5)
	fmt.Printf("%c\n", x6)
}
