package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonStr := `{"name": "John", "age": 30}`

	reader := strings.NewReader(jsonStr)

	decoder := json.NewDecoder(reader)

	var person Person

	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Ошибка чтения JSON:", err)
		return
	}

	fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
}
