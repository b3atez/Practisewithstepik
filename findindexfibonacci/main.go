package main

import (
	"fmt"
)

//0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946,
func main() {
	var n int
	var arr []int
	arr = []int{
		1, 1,
	}
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}

	fmt.Println(arr)

	for i, v := range arr {
		if n == i+1 {
			fmt.Println(v)
			return
		}
	}

}
