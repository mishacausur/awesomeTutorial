package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b rune
	var text string
	fmt.Scanf("%c\n", &a)
	fmt.Scanf("%c\n", &b)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}

	x, y := count(a, b, text)

	if x >= y {
		fmt.Printf("%c %d\n", a, x)
		fmt.Printf("%c %d\n", b, y)
	} else {
		fmt.Printf("%c %d\n", b, y)
		fmt.Printf("%c %d\n", a, x)
	}
}

func count(a, b rune, text string) (int, int) {
	var x, y int
	for _, char := range text {
		if char == a {
			x++
		} else if char == b {
			y++
		}
	}

	return x, y
}
