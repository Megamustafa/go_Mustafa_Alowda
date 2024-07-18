package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)

	binaries := make([]string, num+1)
	for i := 0; i <= num; i++ {
		binaries[i] = strconv.FormatInt(int64(i), 2)
	}

	fmt.Printf("Binary representations from 0 to %d: %v\n", num, binaries)
}
