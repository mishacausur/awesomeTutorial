package main

import "fmt"

func main() {
	output, op := FindMinMaxInSlice([]int{1, 2, 3, 4})
	fmt.Print(output, op)
}

func FindMinMaxInSlice(slice []int) (int, int) {
	var min int
	var max int

	if len(slice) == 0 {
		return 0, 0
	}
	for i := 0; i < len(slice); i++ {
		if min > slice[i] {
			min = slice[i]
		}
		if max < slice[i] {
			max = slice[i]
		}
	}
	return min, max
}
