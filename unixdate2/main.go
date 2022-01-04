package main

import (
	"bufio"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "мин. ", "")
	s = strings.ReplaceAll(s, " сек.", "")
	s = strings.Replace(s, " ", ":", -1)
	t, _ := time.Parse("04:05", s)
	d, _ := time.ParseDuration(s)
	nw := t.Add(d)

}
