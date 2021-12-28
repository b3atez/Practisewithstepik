package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	split(n)

}

func split(n int) {
	var nums []int
	if n >= 10 {
		for n > 0 {
			nums = append(nums, n%10)
			n = n / 10
		}
		sum(nums)
	} else {
		fmt.Println(n)
	}

}

func sum(n []int) {
	var s int
	for _, val := range n {
		s += val
	}
	split(s)
}
