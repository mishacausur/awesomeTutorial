package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var data []map[string]interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]byte)

	for _, obj := range data {
		class, ok := obj["class"].(string)
		if !ok {
			return nil, errors.New("decoding error")
		}

		if existingData, exists := result[class]; exists {

			var existingObjects []map[string]interface{}
			err := json.Unmarshal(existingData, &existingObjects)
			if err != nil {
				return nil, err
			}
			existingObjects = append(existingObjects, obj)
			updatedData, err := json.Marshal(existingObjects)
			if err != nil {
				return nil, err
			}
			result[class] = updatedData
		} else {

			initialData := []map[string]interface{}{obj}
			initialDataJSON, err := json.Marshal(initialData)
			if err != nil {
				return nil, err
			}
			result[class] = initialDataJSON
		}
	}

	return result, nil
}

func main() {
	// Пример использования
	jsonData := []byte(`[
		{"class": "9A", "name": "Ivan"},
		{"class": "9A", "name": "John"},
		{"class": "9B", "name": "Oleg"},
		{"class": "9B", "name": "Maria"}
	]`)

	result, err := splitJSONByClass(jsonData)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Выводим результат
	for class, data := range result {
		fmt.Printf("Класс: %s\nДанные: %s\n", class, data)
	}
}
