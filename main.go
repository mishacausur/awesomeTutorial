package main

import "fmt"

func main() {
	var i, a, b, c, d, e, f uint
	fmt.Scan(&i)
	a = i / 100000
	b = (i % 100000) / 10000
	c = (i % 10000) / 1000
	d = (i % 1000) / 100
	e = (i % 100) / 10
	f = (i % 10)
	switch {
	case (a + b + c) == (d + e + f):
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}

}
