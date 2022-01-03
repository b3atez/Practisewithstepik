package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	n, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}

}

func divide(a int, b int) (int, error) {

	if b == 0 {
		err := errors.New("ошибка")
		return -1, err
	}
	return a / b, nil
}
