package main

import "fmt"

func main() {
	output := FindMaxKey(map[int]int{10: 37, 3: 19})
	fmt.Print(output)
}

func FindMaxKey(m map[int]int) int {
	var max int
	for key := range m {
		max = key
	}
	for key := range m {
		if key > max {
			max = key
		}
	}
	return max
}
