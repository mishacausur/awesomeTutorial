package main

import "fmt"

func main() {
	var number int
	var res = 0
	fmt.Scanln(&number)

	for i := 1; i <= number; i++ {
		if (i%3 == 0) || (i%5 == 0) {

			continue
		}
		res += i
	}
	fmt.Println(res)
}
