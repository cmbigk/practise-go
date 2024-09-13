package main

import "fmt"

func main() {
	for b := 'z'; b >= 'a'; b-- {
		fmt.Printf("%c", b)
	}
	fmt.Println("\n")
}
