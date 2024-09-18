package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file for reading
	inputFile, err := os.Open("haha.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Read the entire content of the input file
	scanner := bufio.NewScanner(inputFile)
	var content []string
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Convert all the text to lowercase
	for i, line := range content {
		content[i] = strings.ToLower(line)
	}

	// Capitalize one word (for example, the word "example")
	wordToCapitalize := "doit"
	for i, line := range content {
		content[i] = strings.ReplaceAll(line, wordToCapitalize, strings.Title(wordToCapitalize))
	}

	// Create the output file for writing
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write the modified content to the output file
	writer := bufio.NewWriter(outputFile)
	for _, line := range content {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	// Flush the buffered data to the output file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing output file:", err)
	}
}
