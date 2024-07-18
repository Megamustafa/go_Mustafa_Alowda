package main

import (
	"fmt"
)

// Function to generate Pascal's Triangle
func generatePascalTriangle(rows int) [][]int {
	triangle := make([][]int, rows)

	for i := 0; i < rows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

// Function to print Pascal's Triangle
func printPascalTriangle(triangle [][]int) {
	for _, row := range triangle {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	var rows int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	triangle := generatePascalTriangle(rows)
	printPascalTriangle(triangle)
}
