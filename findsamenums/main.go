package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Scan(&num1, &num2)
	n1 := split(num1)
	n2 := split(num2)

	for _, i := range n1 { //find same nums
		for _, j := range n2 {
			if i == j {
				fmt.Print(i, " ")
			}
		}
	}

}

func split(num int) []int { //function return number as a massive
	var t []int
	for num > 0 { //create massive from number
		t = append(t, num%10)
		num = num / 10
	}
	var nums []int
	for i := len(t) - 1; i >= 0; i-- { //re crete vice versa
		nums = append(nums, t[i])
	}
	return nums
}
