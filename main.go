package main

import (
	"fmt"
	"strings"
)

func main() {

	var count int
	fmt.Scan(&count)

	people := make([]string, count)

	for i := 0; i < count; i++ {
		fmt.Scan(&people[i])
	}

	for {
		var prefix string
		if _, error := fmt.Scan(&prefix); error != nil {
			break
		}

		result := false
		for _, name := range people {
			if strings.HasPrefix(name, prefix) {
				fmt.Println(name)
				result = true
				break
			}
		}
		if !result {
			fmt.Println("Не найдено")
		}
	}
}
