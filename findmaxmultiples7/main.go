package main

import "fmt"

func main() {
	var a, b, n, c int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		if i%7 == 0 {
			n = i
			c++
		}
	}

	if c != 0 {
		fmt.Println(n)
	} else {
		fmt.Println("NO")
	}

}
