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
	t, _ := time.Parse("2006-01-02 15:04:05", s)

	if t.Format("15:04") > "13:00" {
		fmt.Println(t.Format("2006-01-02 15:04:05"))
	} else {
		nw := t.AddDate(0, 0, 1)
		fmt.Println(nw.Format("2006-01-02 15:04:05"))
	}

}
