package main

import (
	"fmt"
	"sort"
)

func main() {
	output := SortSlice([]int{1, 4, 7, 2})
	fmt.Print(output)
}

func SortSlice(slice []int) []int {
	sort.Sort(sort.IntSlice(slice))
	return slice
}
