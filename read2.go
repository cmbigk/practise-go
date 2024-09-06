package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("haha.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read each line and print it
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check for any errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
