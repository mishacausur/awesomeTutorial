package main

import "fmt"

func main() {
	output := StringLength("evrb")
	fmt.Print(output)
}

func StringLength(input string) int {
	return len(input)
}
