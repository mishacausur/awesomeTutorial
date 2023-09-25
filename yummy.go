package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	var result = make(map[string]string)

	for key := range m {
		result[m[key]] = key
	}

	return result
}
