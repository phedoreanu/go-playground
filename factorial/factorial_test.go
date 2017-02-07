package factorial_test

import (
	"fmt"
	"testing"

	"github.com/phedoreanu/go-playground/factorial"
)

const LIMIT = 65

func Example() {
	fmt.Println(factorial.Solution(0))
	// Output: 1
}

func Example1() {
	fmt.Println(factorial.Solution(1))
	// Output: 1
}

func Example2() {
	fmt.Println(factorial.Solution(2))
	// Output: 2
}

func Example4() {
	fmt.Println(factorial.Solution(4))
	// Output: 24
}

func Example5() {
	fmt.Println(factorial.Solution(5))
	// Output: 120
}

func Example10() {
	fmt.Println(factorial.Solution(10))
	// Output: 3628800
}

func Example13() {
	fmt.Println(factorial.Solution(13))
	// Output: 6227020800
}

func Example20() {
	fmt.Println(factorial.Solution(20))
	// Output: 2432902008176640000
}

func Example21() {
	fmt.Println(factorial.Solution(21))
	// Output: 14197454024290336768
}

func Example25() {
	fmt.Println(factorial.Solution(25))
	// Output: 7034535277573963776
}

func Example50() {
	fmt.Println(factorial.Solution(50))
	// Output: 15188249005818642432
}

func Example60() {
	fmt.Println(factorial.Solution(60))
	// Output: 9727775195120271360
}

func Example65() {
	fmt.Println(factorial.Solution(65))
	// Output: 9223372036854775808
}

func Example66() {
	fmt.Println(factorial.Solution(66))
	// Output: 0
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.Loop(LIMIT)
	}
}

func BenchmarkRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.Recursion(LIMIT)
	}
}

func BenchmarkTailRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.TailRecursion(LIMIT, 1)
	}
}
