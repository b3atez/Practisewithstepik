package main

import "fmt"

func main() {
	var n, num, sum int
	var nums []int
	{
	}

	fmt.Scan(&n)

	for i := 0; i < n; i++ { //get numbers
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	for _, i := range nums { //counting sum of number which multiples of 8
		if i >= 10 && i < 100 {
			if i%8 == 0 {
				sum += i
			}
		}
	}

	fmt.Println(sum) // print sum
}
