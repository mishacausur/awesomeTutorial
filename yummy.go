package main

import (
	"fmt"
	"math"
)

func main() {
	SqRoots()
}

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	var result = b*b - 4*a*c
	if result > 0 {
		var firstRoot = (-b + math.Sqrt(result)) / (2 * a)
		var secondRoot = (-b - math.Sqrt(result)) / (2 * a)
		fmt.Println(math.Min(firstRoot, secondRoot), math.Max(firstRoot, secondRoot))
	} else if result == 0 {
		var root = -b / (2 * a)
		fmt.Println(root)
	} else {
		fmt.Println(0, 0)
	}
}
