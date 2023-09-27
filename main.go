package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       string `json:"-"`
	privateNotes string
}

func main() {
	jsonStr := `{"name": "John", "age": 30, "Gender": "male"}`
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
}
