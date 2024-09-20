package main

import (
	"fmt"
	"strconv"
)

func BasicAtoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}
func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	fmt.Println(BasicAtoi(""))

	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000012345"))
	fmt.Println(BasicAtoi("000000"))
}
