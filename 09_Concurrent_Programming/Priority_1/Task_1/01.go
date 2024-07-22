package main

import (
	"fmt"
	"time"
)

func reverse(word string, ch chan string) {
	length := len(word)
	for i := 1; i <= length; i++ {
		ch <- reverseSubstr(word[:i])
		time.Sleep(3 * time.Second)
	}
	close(ch)
}

func reverseSubstr(substr string) string {
	reversed := ""
	for _, char := range substr {
		reversed = string(char) + reversed
	}
	return reversed
}

func main() {
	word := "phone"
	ch := make(chan string)
	go reverse(word, ch)
	for reversed := range ch {
		fmt.Println(reversed)
	}
}
