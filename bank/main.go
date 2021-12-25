package main

import "fmt"

func main() {
	var x, p, y float64
	var count int

	fmt.Scan(&x, &p, &y)

	for y > x { //calculating
		x += x * p / 100 //find percent p
		n := int(x)      //remove nums after dot
		x = float64(n)
		count++ //counting year

	}

	fmt.Println(count)

}
