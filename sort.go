package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
}
