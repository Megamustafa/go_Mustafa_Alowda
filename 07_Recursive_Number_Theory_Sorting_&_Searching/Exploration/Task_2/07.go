package main

import (
	"fmt"
)

func main() {
	fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}

func getItem(items []int, target int) []int {
	var result []int
	findCombination(items, target, 0, []int{}, &result)
	return result
}

func findCombination(items []int, target int, start int, currentCombination []int, result *[]int) bool {
	sum := 0
	for _, v := range currentCombination {
		sum += v
	}

	if sum == target {
		*result = append([]int{}, currentCombination...)
	}

	if sum > target {
		return false
	}

	for i := start; i < len(items); i++ {
		currentCombination = append(currentCombination, items[i])
		if findCombination(items, target, i+1, currentCombination, result) {
			return true
		}
		currentCombination = currentCombination[:len(currentCombination)-1]
	}

	return false
}
