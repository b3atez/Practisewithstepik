package main

import "fmt"

func main() {
	var n, t int
	var nums []int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&t)
		nums = append(nums, t)
	}

	fmt.Println(nums[3])
}
