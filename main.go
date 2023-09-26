package main

import (
	"fmt"
	"strings"
)

func main() {
	var bloom, stem int
	fmt.Scanln(&bloom)
	fmt.Scanln(&stem)
	draw(bloom, stem)
}

func draw(bloom, stem int) {

	for i := bloom; i > 0; i-- {
		fmt.Println(strings.Repeat("*", i))
	}

	for i := stem - 1; i > 0; i-- {
		fmt.Println("*")
	}
}
