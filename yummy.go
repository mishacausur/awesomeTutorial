package main

import "fmt"

func main() {
	output := ConcatenateStrings("evrb", "rvrv")
	fmt.Print(output)
}

func ConcatenateStrings(str1, str2 string) string {
	return str1 + " " + str2
}
