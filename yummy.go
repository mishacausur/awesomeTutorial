package main

import "fmt"

func main() {
	output := SumOfSlice([]int{1, 2, 3, 4})
	fmt.Print(output)
}

func SumOfSlice(slice []int) int {
	var result int
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	return result
}
