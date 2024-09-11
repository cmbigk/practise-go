package main

import "fmt"

func main() {
	for g := 'a'; g <= 'z'; g++ {
		fmt.Printf("%c", g)
	}
	fmt.Printf("\n")
}
