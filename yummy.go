package main

import "fmt"

func main() {
	IsPowerOfTwoRecursive(18)
}

func IsPowerOfTwoRecursive(N int) {
	if N == 1 {
		fmt.Print("YES")
	} else if N > 1 && N%2 == 0 {
		IsPowerOfTwoRecursive(N / 2)
	} else {
		fmt.Print("NO")
	}
}
