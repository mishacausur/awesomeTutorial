package main

import "fmt"

func main() {
	var workArray = [10]uint8{}
	var i uint8
	var index = 0

	for fmt.Scan(&i); index <= len(workArray)-1; fmt.Scan(&i) {
		workArray[index] = i
		fmt.Println(workArray)
		index++
	}

	//var indexed = [6]uint{}
	var indexes = [6]uint8{}
	var num uint8
	var index1 = 0
	for fmt.Scan(&num); index1 <= len(indexes)-1; fmt.Scan(&num) {
		fmt.Println(indexes)
		indexes[index1] = num
		index1++
		fmt.Println(indexes)
	}

	var a, b, c, d, e, f uint8
	a = workArray[indexes[0]]
	b = workArray[indexes[1]]
	c = workArray[indexes[2]]
	d = workArray[indexes[3]]
	e = workArray[indexes[4]]
	f = workArray[indexes[5]]

	workArray[indexes[0]] = b
	workArray[indexes[1]] = a
	workArray[indexes[2]] = d
	workArray[indexes[3]] = c
	workArray[indexes[4]] = f
	workArray[indexes[5]] = e

	fmt.Println(workArray)
}
