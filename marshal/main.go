package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type Question struct {
	Domain  string `json:"domain"`
	TypeDNS string `json:"typeDNS"`
}

func main() {
	q := Question{"_dmarc.dina.io.", "TXT"}
	byt, err := json.Marshal(q)

	if err != nil {
		log.Println(err)
	}

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	fmt.Println(string(byt))

	err = ioutil.WriteFile("temp.json", []byte(byt), 0644)

	if err != nil {
		log.Println(err)
	}

	file, err := os.OpenFile("temp.json", os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Println(reflect.TypeOf(file))

	defer file.Close()

	//Print the contents of the file
	data, err := ioutil.ReadFile("temp.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
