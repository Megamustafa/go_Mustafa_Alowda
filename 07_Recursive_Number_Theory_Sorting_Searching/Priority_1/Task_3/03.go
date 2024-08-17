package main

import "fmt"

func main() {
	fmt.Println(getSequence(4))   // 10
	fmt.Println(getSequence(15))  // 120
	fmt.Println(getSequence(100)) // 5050
}

func getSequence(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	sequence := []int{0, 1}

	for i := 2; i <= n; i++ {
		nextNumber := sequence[i-1] + i
		sequence = append(sequence, nextNumber)
	}

	return sequence[n]
}
