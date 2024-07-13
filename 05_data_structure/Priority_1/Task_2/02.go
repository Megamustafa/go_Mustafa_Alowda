package main

import "fmt"

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30}))
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))
	fmt.Println(primeSum([]int{15, 12, 9, 10}))
}

func primeSum(numbers []int) int {
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		if n == 2 {
			return true
		}
		if n%2 == 0 {
			return false
		}
		for i := 3; i*i <= n; i += 2 {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	sum := 0
	for _, number := range numbers {
		if isPrime(number) {
			sum += number
		}
	}
	return sum

}
