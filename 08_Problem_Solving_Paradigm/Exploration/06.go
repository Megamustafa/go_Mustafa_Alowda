package main

import (
	"fmt"
)

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanNum := ""
	i := 0

	for num > 0 {
		for num >= val[i] {
			romanNum += syb[i]
			num -= val[i]
		}
		i++
	}
	return romanNum
}

func main() {
	// Test cases
	fmt.Println(intToRoman(4))    // Output: IV
	fmt.Println(intToRoman(9))    // Output: IX
	fmt.Println(intToRoman(23))   // Output: XXIII
	fmt.Println(intToRoman(2021)) // Output: MMXXI
	fmt.Println(intToRoman(1646)) // Output: MDCXLVI
}
