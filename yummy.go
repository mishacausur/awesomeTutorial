package main

import "fmt"

func main() {
	var number int
	var counter = 0
	fmt.Scanln(&number)

	for i := 2; counter < 10; i++ {
		var res = fib(i)
		if res < number {
			continue
		}
		fmt.Println(res)
		counter++
	}
}

func fib(n int) int {

	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
