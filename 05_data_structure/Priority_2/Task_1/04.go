package main

import "fmt"

func main() {
	fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))
	fmt.Println(spinSlice([]int{6, 7, 8}))
	fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1}))
}

func spinSlice(numbers []int) []int {
	length := len(numbers)
	if length == 0 {
		return numbers
	}

	result := make([]int, length)
	for i := 0; i < (length+1)/2; i++ {
		result[i*2] = numbers[i]
		if i*2+1 < length {
			result[i*2+1] = numbers[length-1-i]
		}
	}

	return result
}
