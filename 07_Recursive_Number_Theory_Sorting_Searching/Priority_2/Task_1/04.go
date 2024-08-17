package main

import (
	"fmt"
)

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	catalanNumbers := make([]int, n+1)
	catalanNumbers[0] = 1
	catalanNumbers[1] = 1

	for i := 2; i <= n; i++ {
		catalanNumbers[i] = 0
		for j := 0; j < i; j++ {
			catalanNumbers[i] += catalanNumbers[j] * catalanNumbers[i-1-j]
		}
	}

	return catalanNumbers[n]
}
