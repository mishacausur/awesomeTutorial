package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func main() {
	inputJSON := []byte(`[
			{
				"name": "Oleg",
				"grade": 12
			}
		]`)

	expectedJSON := []byte(`[{"name":"Oleg","grade":13}]`)

	updatedJSON, err := modifyJSON(inputJSON)
	if err != nil {
		fmt.Println("Error while modifying JSON: %v", err)
	}

	if !bytes.Equal(updatedJSON, expectedJSON) {
		fmt.Println("Expected updated JSON to be %s, but got %s", expectedJSON, updatedJSON)
	}
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	students := []Student{}
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		fmt.Println("Ошибка при чтении JSON-данных:", err)
		return jsonData, err
	}
	for i := range students {
		students[i].Grade += 1
	}

	updatedJSON, err := json.Marshal(students)
	if err != nil {
		fmt.Println("Ошибка при записи JSON-данных:", err)
		return jsonData, err
	}

	return updatedJSON, nil
}
