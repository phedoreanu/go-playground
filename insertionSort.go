package main

import (
	"fmt"
)

func insertionSort(a []int, asc bool) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			switch asc {
			case true:
				if a[j] < a[j-1] {
					a[j], a[j-1] = a[j-1], a[j]
				}
			case false:
				if a[j] > a[j-1] {
					a[j], a[j-1] = a[j-1], a[j]
				}
			}
		}
	}

}

func main() {
	a := []int{5, 2, 6, 3, 1, 4}
	b := []int{31, 41, 59, 26, 41, 58}
	insertionSort(a, true)
	fmt.Println(a)
	insertionSort(a, false)
	fmt.Println(a)

	fmt.Println()

	insertionSort(b, true)
	fmt.Println(b)
	insertionSort(b, false)
	fmt.Println(b)
}
