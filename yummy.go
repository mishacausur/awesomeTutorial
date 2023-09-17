package main

import "fmt"

func main() {
	var number int
	var result = 0
	fmt.Scanln(&number)
	var text = "Некорректный ввод"
	for i := 1; i <= number; i++ {
		if number < 0 {
			break
		}
		if i%2 != 0 {
			result = result + i
			text = fmt.Sprintln(result)
		}
	}
	fmt.Println(text)
}
