package main

import "fmt"

func main() {
	var number int
	var result = 0
	fmt.Scanln(&number)
	for i := 1; i <= number; i++ {
		result = result + i
	}
	fmt.Println(result)
}
