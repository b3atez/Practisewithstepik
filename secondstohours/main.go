package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Printf("It is %d hours %d minutes.", num/3600, (num%3600)/60) //1 hour = 3600 seconds; 1 minute = 60 seconds

}
