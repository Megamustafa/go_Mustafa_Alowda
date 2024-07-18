package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	if n == 0 {
		return 0
	}

	// Create a dp array to store the minimum cost to reach each rock
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		if i+1 < n {
			dp[i+1] = min(dp[i+1], dp[i]+int(math.Abs(float64(jumps[i+1]-jumps[i]))))
		}
		if i+2 < n {
			dp[i+2] = min(dp[i+2], dp[i]+int(math.Abs(float64(jumps[i+2]-jumps[i]))))
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
