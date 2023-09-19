package main

func main() {
	ReverseSlice([]int{1, 2, 3, 4})
}

func ReverseSlice(slice []int) []int {
	var reversed = make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		reversed[i] = slice[len(slice)-1-i]
	}
	return reversed
}
