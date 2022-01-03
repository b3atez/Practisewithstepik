package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	s := strings.Split(str, "")
	fmt.Println(strings.Join(s, "*"))
}
