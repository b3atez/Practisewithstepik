package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.ReplaceAll(s, "\n", "")
	lst := strings.Split(s, ",")
	t1, _ := time.Parse("02.01.2006 15:04:05", lst[0])
	t2, _ := time.Parse("02.01.2006 15:04:05", lst[1])

	var diff time.Duration

	if t1.Before(t2) {
		diff = t2.Sub(t1)

	} else {
		diff = t1.Sub(t2)

	}
	fmt.Println(diff)
}
