package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%10 == 1 {
		fmt.Printf("%d korova", n)
	} else if n%10 >= 1 && n%10 < 5 && n >= 12 && n <= 14 {
		fmt.Printf("%d korovy", n)
	} else if n%10 >= 5 && n%10 == 0 {
		fmt.Printf("%d korov", n)
	}

}
