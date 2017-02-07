package main

import "fmt"

func main() {
	s := "Let's hack together!"
	fmt.Printf("%# x\n", s)
	fmt.Println([]byte(s))
}
