package main

import "fmt"

func main() {
	var n, c, t int
	var nums []int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&t)
		nums = append(nums, t)
	}
	min := nums[0]
	for _, val := range nums {
		if min > val {
			min = val
		}
	}

	for _, val := range nums {
		if val == min {
			c++
		}
	}

	fmt.Println(c)
}
