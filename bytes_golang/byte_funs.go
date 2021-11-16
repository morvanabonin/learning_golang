// Tests based on this article
// https://zetcode.com/golang/byte/

package main

import (
	"bytes"
	"fmt"
)

func main() {

	data1 := []byte{102, 97, 108, 99, 111, 110} //falcon
	data2 := []byte{111, 110}                   // on

	if bytes.Contains(data1, data2) {
		fmt.Println("contains")
	} else {
		fmt.Println("does not contain")
	}

	if bytes.Equal([]byte("falcon"), []byte("owl")) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	data3 := []byte{111, 119, 108, 9, 99, 97, 116, 32, 32, 32, 32, 100, 111, 103, 32, 112, 105, 103, 32, 32, 32, 32, 98, 101, 97, 114}
	// data3 := []byte{131, 166, 65, 99, 116, 105, 111, 110, 206, 0, 43, 101, 255, 167, 67, 111, 110, 116, 97, 99, 116, 207, 2, 50, 166, 220, 237, 125, 109, 255, 101, 242, 121, 183, 165, 244, 36, 0}
	// data3 := []byte{131, 166, 65, 99, 116, 105, 111, 110, 206, 0, 23, 143, 58, 167, 67, 111, 110, 116, 97, 99, 116, 206, 0, 1, 12, 181, 164, 84, 105, 109, 101, 210, 93, 39, 137, 122}

	fields := bytes.Fields(data3)

	fmt.Println(fields)

	for _, e := range fields {
		fmt.Printf("%s\n", string(e))
	}
}
