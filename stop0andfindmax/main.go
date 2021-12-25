package main

import "fmt"

func main() {
	var num, counter int
	var nums []int
	{
	}

	for { //input values until entered 0
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		nums = append(nums, num) // append values
	}

	max := nums[0]

	for _, i := range nums { //find max value
		if max < i {
			max = i
		}
	}
	for _, i := range nums { //counting max values
		if max == i {
			counter++
		}
	}
	fmt.Println(counter)

}
