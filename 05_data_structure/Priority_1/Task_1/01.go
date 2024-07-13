package main

import "fmt"

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1)

	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2)

	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3)
}

func merge(data [][]int) []int {
	uniqueElements := make(map[int]bool)
	var result []int

	for _, row := range data {
		for _, element := range row {
			if _, exists := uniqueElements[element]; !exists {
				uniqueElements[element] = true
				result = append(result, element)
			}
		}
	}

	return result

}
