package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&num, &n)
	arr := remove(n, split(num))
	for _, v := range arr {
		fmt.Print(v)
	}
	fmt.Println()
}

func remove(n int, arr []int) []int {
	var nw []int
	for _, v := range arr {
		if n != v {
			nw = append(nw, v)
		}
	}
	return nw
}

func split(num int) []int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10

	}

	var nw []int

	for i := len(arr) - 1; i >= 0; i-- {
		nw = append(nw, arr[i])
	}
	return nw
}
