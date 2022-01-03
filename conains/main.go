package main

import (
	"fmt"
	"strings"
)

func main() {
	var text, sym string
	fmt.Scan(&text, &sym)

	if strings.Contains(text, sym) {
		fmt.Println(strings.Index(text, string(sym[0])))
		return
	}

	fmt.Println(-1)
}
