package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Printf("%d korova", n)
	} else if n >= 2 && n < 5 {
		fmt.Printf("%d korovy", n)
	} else if n >= 5 {
		fmt.Printf("%d korov", n)
	}
}
