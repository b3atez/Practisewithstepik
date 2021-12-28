package main

import (
	"fmt"
)

func main() {
	var n, t, c int
	var nums []int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&t)
		nums = append(nums, t)
	}

	for _, val := range nums {
		if val >= 1 {
			c++
		}
	}

	fmt.Println(c)

}
