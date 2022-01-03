package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	if isUp(text) && finddot(text) {
		fmt.Println("Right")
		return
	}
	fmt.Println("Wrong")

}

func isUp(t string) bool {

	if isRussianUpper(t) {
		return true
	}

	return false
}

func finddot(t string) bool {
	for _, v := range t {
		if string(v) == "." {
			return true
		}
	}
	return false
}

func isRussianUpper(text string) bool {
	for _, r := range []rune(text) {
		if r < '\u0410' || r > '\u042F' {
			return false
		}
	}
	return true
}
