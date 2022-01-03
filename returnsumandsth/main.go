package main

import "fmt"

func main() {
	a, b := sumInt(1, 2)
	fmt.Println(a, b)
}

func sumInt(a ...int) (int, int) {
	var sum int
	for _, v := range a {
		sum += v
	}
	return len(a), sum

}
