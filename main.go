package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var person Person
	err = json.Unmarshal(byteValue, &person)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	fmt.Println("Имя:", person.Name)
	fmt.Println("Email:", person.Email)
}
