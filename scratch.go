package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	var hours = d / 30
	var minutes = (d % 30) * 2
	fmt.Print("It is ", hours, " hours, ", minutes, " minutes")
}
