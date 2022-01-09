package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	switch {
	case i > 0:
		fmt.Print("Число положительное")
	case i < 0:
		fmt.Print("Число отрицательное")
	default:
		fmt.Print("Ноль")
	}
}
