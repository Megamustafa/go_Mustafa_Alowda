package main

import "fmt"

func primeNumber(number int) bool {

	for i := 2; i < number; i++ {

		if number%i == 0 {

			fmt.Println("The number: ", number, "is not a prime number")
			return false

		}

	}
	fmt.Println("The number: ", number, "is a prime number")
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))

}
