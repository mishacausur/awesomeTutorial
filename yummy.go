package main

import "fmt"

func main() {
	output := IntersectionOfSlices([]int{1, 2}, []int{1, 3, 4})
	fmt.Print(output)
}

func IntersectionOfSlices(slice1, slice2 []int) []int {
	var result = make([]int, 0)
	for i := 0; i < len(slice1); i++ {
		for y := 0; y < len(slice2); y++ {
			if slice1[i] == slice2[y] {
				result = append(result, slice1[i])
			}
		}
	}
	return result
}
