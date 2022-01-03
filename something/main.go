package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	lst := split(num)
	for _, v := range lst {
		fmt.Print(v * v)
	}
}

func split(num int) []int {
	var lst, nw []int
	for num > 0 {
		lst = append(lst, num%10)
		num = num / 10
	}
	for i := len(lst) - 1; i >= 0; i-- {
		nw = append(nw, lst[i])
	}

	return nw
}
