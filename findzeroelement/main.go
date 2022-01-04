package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dataFromFile, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
		return
	}
	lst := strings.Split(string(dataFromFile), ";")

	for i, v := range lst {
		if v == "0" {
			fmt.Println(i + 1)
			return
		}
	}
}
