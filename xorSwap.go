package main

import (
	"fmt"
)

func xorSwap(x, y *int) {
	if x != y {
		*x ^= *y
		*y ^= *x
		*x ^= *y
	}
}

func main() {
	x, y := 1, 0
	fmt.Println(x, y)
	xorSwap(&x, &y)
	fmt.Println(x, y)
}
