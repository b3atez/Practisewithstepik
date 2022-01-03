package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var c int
	fmt.Scan(&str)

	for i := 0; i < len(str); i++ {
		for j := 1; j < len(str); j++ {
			if str[i] == str[j] {
				c++
			}
		}
		if c > 1 {
			str = strings.ReplaceAll(str, string(str[i]), "")
		}
		c = 0
	}

	fmt.Println(str)

}
