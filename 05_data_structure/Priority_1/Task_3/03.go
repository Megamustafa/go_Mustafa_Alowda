package main

import "fmt"

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Println("sum: ", sum(data1))
	fmt.Println("mean: ", mean(data1))
	fmt.Println("median: ", median(data1))
	fmt.Println("mode: ", mode(data1))

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Println("sum: ", sum(data2))
	fmt.Println("mean: ", mean(data2))
	fmt.Println("median: ", median(data2))
	fmt.Println("mode: ", mode(data2))
}

func sum(data []float64) float64 {
	sum := 0.0
	for _, number := range data {
		sum += number
	}
	return sum
}

func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := sum(data)
	return sum / float64(len(data))
}

func median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sortedNumbers := make([]float64, len(data))
	copy(sortedNumbers, data)
	n := len(sortedNumbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortedNumbers[j] > sortedNumbers[j+1] {
				sortedNumbers[j], sortedNumbers[j+1] = sortedNumbers[j+1], sortedNumbers[j]
			}
		}
	}

	middle := n / 2
	if n%2 == 0 {
		return (sortedNumbers[middle-1] + sortedNumbers[middle]) / 2
	}
	return sortedNumbers[middle]
}

func mode(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	frequency := make(map[float64]int)
	for _, number := range data {
		frequency[number]++
	}

	maxCount := 0
	var mode float64
	for number, count := range frequency {
		if count > maxCount {
			maxCount = count
			mode = number
		}
	}
	return mode
}
