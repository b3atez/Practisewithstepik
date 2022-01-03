package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	lst := strings.Split(str, "")
	max := lst[0]

	for _, v := range lst {
		if max < v {
			max = v
		}
	}

	fmt.Println(max)
}
