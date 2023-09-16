package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	for i := 1; i <= 10; i++ {
		var result = number * i
		fmt.Println(number, "*", i, "=", result)
	}
}
