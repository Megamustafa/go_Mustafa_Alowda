package main

import "fmt"

func fibonacci(number int) int {

	table := []int{0, 1}

	for i := 2; i <= number; i++ {
		table = append(table, table[i-1]+table[i-2])
	}

	return table[number]
}

func main() {

	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144

}
