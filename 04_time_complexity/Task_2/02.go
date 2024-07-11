package main

import (
	"fmt"
)

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}

func main() {

	fmt.Println(pow(2, 3))

	fmt.Println(pow(5, 3))

	fmt.Println(pow(10, 2))

	fmt.Println(pow(2, 5))

	fmt.Println(pow(7, 3))

}
