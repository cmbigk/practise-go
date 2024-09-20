package main

import (
	"fmt"
)

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		}
		result = (result*10 + int(char-'0'))
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
