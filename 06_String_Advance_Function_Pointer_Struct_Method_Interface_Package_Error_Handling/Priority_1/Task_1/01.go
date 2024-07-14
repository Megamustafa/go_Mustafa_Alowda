package main

import "fmt"

func main() {
	fmt.Println(pickVocals("alterra academy"))
	fmt.Println(pickVocals("i love programming"))
	fmt.Println(pickVocals("go is awesome programming language"))
}

func pickVocals(sentence string) string {

	vowels := "aeiouAEIOU"
	var result string

	for _, char := range sentence {
		if char == ' ' {
			result += " "
		} else {
			for _, v := range vowels {
				if char == v {
					result += string(char)
					break
				}
			}
		}
	}
	return result

}
