package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var t []int
	for i := 1; i <= n; {
		t = append(t, i)
		i = i * 2
	}

	code := rewrite(t)
	fmt.Println(coding(n, code))

}

func coding(n int, code []int) int {
	var b int

	for _, v := range code {
		b = b * 10
		if n-v >= 0 {
			n = n - v
			b += 1
		}
	}

	return b
}

func rewrite(c []int) []int {
	var nw []int
	for i := len(c) - 1; i >= 0; i-- {
		nw = append(nw, c[i])
	}
	return nw
}
