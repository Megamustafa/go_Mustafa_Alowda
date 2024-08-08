package main

import (
	"net/http"
	"strings"
	"unicode"

	"github.com/labstack/echo/v4"
)

// Word represents the word data structure
type Word struct {
	Word        string `json:"word"`
	Length      int    `json:"length"`
	NumOfVocals int    `json:"num_of_vocals"`
	Palindrome  bool   `json:"palindrome"`
}

// In-memory storage for words
var words []Word

func main() {
	e := echo.New()

	// Routes
	e.GET("/words", getWords)
	e.POST("/words", addWord)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// getWords handles the GET request to fetch all word data
func getWords(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    words,
		"message": "all words",
	})
}

// addWord handles the POST request to add a new word
func addWord(c echo.Context) error {
	wordRequest := new(struct {
		Word string `json:"word"`
	})
	if err := c.Bind(wordRequest); err != nil || strings.TrimSpace(wordRequest.Word) == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed to add word data",
		})
	}

	wordData := Word{
		Word:        wordRequest.Word,
		Length:      len([]rune(wordRequest.Word)),
		NumOfVocals: countVocals(wordRequest.Word),
		Palindrome:  isPalindrome(wordRequest.Word),
	}

	words = append(words, wordData)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "word added",
		"data":    wordData,
	})
}

// countVocals counts the number of vowels in a word
func countVocals(word string) int {
	splitted := strings.Split(word, "")
	count := 0
	for i := 0; i < len(splitted); i++ {
		if splitted[i] == "a" || splitted[i] == "e" || splitted[i] == "i" || splitted[i] == "o" || splitted[i] == "u" || splitted[i] == "A" || splitted[i] == "E" || splitted[i] == "I" || splitted[i] == "O" || splitted[i] == "U" {
			count++
		}
	}
	return count
}

// isPalindrome checks if a word is a palindrome
func isPalindrome(word string) bool {
	// splitted := strings.Split(word, "")
	// reversed := []string{}
	// for i := len(splitted) - 1; i >= 0; i-- {
	// 	reversed = append(reversed, splitted[i])
	// }
	// joined := strings.Join(reversed, "")

	// return word == joined
	normalized := strings.ToLower(strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, word))

	runes := []rune(normalized)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}
