package main

import "fmt"

func main() {
	var i, sum int
	fmt.Scan(&i)
	for i != 0 {
		var y int
		fmt.Scan(&y)
		if (y > 10 && y < 99) && (y%8 == 0) {
			sum += y
		}
		i--
	}
	fmt.Print(sum)
}
