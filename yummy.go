package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	for i := 1; i <= number; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
