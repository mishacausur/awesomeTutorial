package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y)
	for a := 1000; x/a >= 0; a = a % 10 {
		if x/a == 0 {
			continue
		}
		z = (x / a) % 10
		for e := 1000; y/e >= 0; e = e % 10 {
			if z == (y/e)%10 {
				fmt.Print(z, " ")
			}
			if e == 1 {
				break
			}
		}
		if a == 1 {
			break
		}
	}
}
