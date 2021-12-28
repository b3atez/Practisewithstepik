package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < n; {
		fmt.Printf("%d ", i)
		i = i * 2
	}
}
