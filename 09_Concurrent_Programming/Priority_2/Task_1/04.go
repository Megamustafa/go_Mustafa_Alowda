package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

func main() {
	// Open the text file
	file, err := os.Open("/mnt/data/document.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a channel to pass word frequencies
	wordChan := make(chan map[string]int)
	var wg sync.WaitGroup

	// Split the text into words
	re := regexp.MustCompile(`[a-zA-Z]+`)
	words := re.FindAllString(text, -1)

	// Divide words into chunks and process each chunk in a separate goroutine
	chunkSize := len(words) / 4
	for i := 0; i < len(words); i += chunkSize {
		wg.Add(1)
		end := i + chunkSize
		if end > len(words) {
			end = len(words)
		}
		go func(chunk []string) {
			defer wg.Done()
			wordFreq := make(map[string]int)
			for _, word := range chunk {
				word = strings.ToLower(word)
				wordFreq[word]++
			}
			wordChan <- wordFreq
		}(words[i:end])
	}

	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(wordChan)
	}()

	// Collect word frequencies from all goroutines
	totalWordFreq := make(map[string]int)
	for wf := range wordChan {
		for word, freq := range wf {
			totalWordFreq[word] += freq
		}
	}

	// Create a CSV file to store the word frequencies
	outputFile, err := os.Create("word_frequencies.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"word", "frequencies"})

	// Write word frequencies
	for word, freq := range totalWordFreq {
		writer.Write([]string{word, fmt.Sprintf("%d", freq)})
	}

	fmt.Println("Word frequencies have been written to word_frequencies.csv")
}
