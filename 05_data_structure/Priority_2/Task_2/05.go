package main

import "fmt"

func main() {
	fmt.Println(groupPrime([]int{2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(groupPrime([]int{15, 24, 3, 9, 5}))
	fmt.Println(groupPrime([]int{4, 8, 9, 12}))
}

func groupPrime(numbers []int) []any {

	var nonPrimes []int
	var primes []int

	isPrime := func(num int) bool {
		if num <= 1 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	for _, num := range numbers {
		if isPrime(num) {
			primes = append(primes, num)
		} else {
			nonPrimes = append(nonPrimes, num)
		}
	}

	var result []any

	if len(primes) > 0 {
		primesStr := "["
		for i, num := range primes {
			if i > 0 {
				primesStr += ", "
			}
			primesStr += fmt.Sprintf("%d", num)
		}
		primesStr += "]"
		result = append(result, primesStr)
	}

	for _, num := range nonPrimes {
		result = append(result, num)
	}

	return result

}
