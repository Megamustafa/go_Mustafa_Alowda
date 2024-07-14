package main

import "fmt"

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
	n := len(sentence)
	if n < 2 {
		return sentence
	}
	return sentence[:1] + sentence[n-1:n] + sentence[1:n-1]
}
