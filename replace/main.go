package main

import "fmt"

//99 151 137 71 117 187 20 93 187 67 1 2 3 5 7 8
func main() {
	var workArray = [10]uint8{}
	var n uint8

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		workArray[i] = n
	}

	for i := 0; i < 3; i++ {
		var a, b uint8
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for _, v := range workArray {
		fmt.Printf("%d ", v)
	}

}
