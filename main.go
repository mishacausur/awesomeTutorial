package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Student struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

package main

import (
"encoding/json"
)

type Student struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	mergedData := []Student{}
	for _, jsonData := range jsonDataList {
		data := []Student{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			return nil, err
		}
		mergedData = append(mergedData, data...)
	}
	mergedJSON, err := json.Marshal(mergedData)
	if err != nil {
		return nil, err
	}
	return mergedJSON, nil
}

func main() {
	inputJSON1 := []byte(`[
			{
				"name": "Oleg",
				"class": "9B"
			},
			{
				"name": "Ivan",
				"class": "9A"
			}
		]`)

	inputJSON2 := []byte(`[
			{
				"name": "Maria",
				"class": "9B"
			},
			{
				"name": "John",
				"class": "9A"
			}
		]`)

	expectedJSON := []byte(`[{"class":"9B","name":"Oleg"},{"class":"9A","name":"Ivan"},{"class":"9B","name":"Maria"},{"class":"9A","name":"John"}]`)

	mergedJSON, err := mergeJSONData(inputJSON1, inputJSON2)
	if err != nil {
		fmt.Println("Error while merging JSON data: %v", err)
	}

	if !bytes.Equal(mergedJSON, expectedJSON) {
		fmt.Println("Expected merged JSON data to be %s, but got %s", expectedJSON, mergedJSON)
	}

	fmt.Println(string(mergedJSON))
}
