package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}

func fibX(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	sum, a, b := 1, 0, 1
	for i := 2; i <= n; i++ {
		next := a + b
		sum += next
		a, b = b, next
	}

	return sum
}
