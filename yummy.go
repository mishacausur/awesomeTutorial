package main

import "fmt"

func main() {
	output := SumOfValuesInMap(map[int]int{10: 37, 3: 19})
	fmt.Print(output)
}

func SumOfValuesInMap(m map[int]int) int {
	var result int

	for key := range m {
		result += m[key]
	}
	return result
}
