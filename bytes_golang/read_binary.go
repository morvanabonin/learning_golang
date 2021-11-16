// Tests based on this article
// https://zetcode.com/golang/byte/

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("bytes.bin")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)

		if err != io.EOF {
			fmt.Println((err))
		}
		break
	}

	fmt.Printf("%s", hex.Dump(buf))
}
