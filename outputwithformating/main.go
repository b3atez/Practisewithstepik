package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	if num <= 0 { //if number equal or less from 0 out put 2nums before and 2 nums after dot
		fmt.Printf("число %2.2f не подходит", num)
	} else if num > 10000 { //if num more then 10000 out put with exponential notation
		fmt.Printf("%e", num)
	} else { // else out put numbers quadrate and 4 numbers after dot
		fmt.Printf("%.4f", num*num)
	}

}
