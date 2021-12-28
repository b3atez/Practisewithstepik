package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max := array[0]

	for i, val := range array {
		if max < val {
			max = array[i]
		}
	}
	fmt.Println(max)
}
