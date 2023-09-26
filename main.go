package main

import (
	"fmt"
	"math/big"
)

func main() {
	var end int
	fmt.Scanln(&end)
	count(end)
}

func count(end int) {

	for i := 3; i < end; i += 5 {
		if isNumberPrime(int64(i)) {
			fmt.Print("хоп")
		} else {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
}

func isNumberPrime(number int64) bool {
	result := big.NewInt(number)
	return result.ProbablyPrime(20)
}
