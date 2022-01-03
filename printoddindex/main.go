package main

import "fmt"

func main() {
	var text, nwtext string

	fmt.Scan(&text)

	for i, v := range text {
		if i%2 != 0 {
			nwtext += string(v)
		}
	}

	fmt.Println(nwtext)

}