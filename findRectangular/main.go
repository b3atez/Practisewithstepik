package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a*a+b*b == c*c { //squareness condition from Pythagoras
		fmt.Println("Прямоугольный") //Rectangular
	} else {
		fmt.Println("Непрямоугольный") //not Rectangular
	}
}
