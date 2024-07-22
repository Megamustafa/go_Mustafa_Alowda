package main

import (
	"fmt"
	"math"
)

func isComposite(n int) bool {
	if n <= 1 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func generateComposites(ch chan int) {
	for i := 1; i <= 100; i++ {
		if isComposite(i) {
			ch <- i
		}
	}
	close(ch)
}

func calculateSquare(chIn chan int, chOut chan int) {
	for n := range chIn {
		chOut <- n * n
	}
	close(chOut)
}

func checkEvenOdd(ch chan int) {
	for n := range ch {
		if n%2 == 0 {
			fmt.Println("Pong")
		} else {
			fmt.Println("Ping")
		}
	}
}

func main() {
	composites := make(chan int)
	squares := make(chan int)

	go generateComposites(composites)
	go calculateSquare(composites, squares)
	checkEvenOdd(squares)
}
