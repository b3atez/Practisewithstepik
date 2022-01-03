package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))

}

func vote(x int, y int, z int) int {
	var arr []int
	var a, b int
	arr = append(arr, x, y, z)

	for _, v := range arr {
		if v == 0 {
			a++
		}
		if v == 1 {
			b++
		}
	}

	if a > b {
		return 0
	}

	return 1
}
