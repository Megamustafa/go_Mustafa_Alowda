package main

import "fmt"

func main() {
	primeFactors(20)   // 2, 2, 5
	primeFactors(75)   // 3, 5, 5
	primeFactors(1024) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) {
	if n <= 1 {
		return
	}
	for n%2 == 0 {
		fmt.Print("2 ")
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n /= i
		}
	}
	if n > 2 {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

}
