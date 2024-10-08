package main

import (
	"fmt"
	"strings"
)

func learningStrContain() bool { //to check if something is in a string
	str := "Hello, how are you?"
	return strings.Contains(str, "love")
}

func lrnStrHassuffix() bool {
	str := "Hello, how are you?" //to check if a string has a suffix, also we can use HasPrefix to check the prefix
	return strings.HasSuffix(str, "you?")
}

func lrnStrLower() string {
	str := "HELLO, HOW ARE YOU?" //to convert a string to lowercase, we can use ToUpper to convert to uppercase
	return strings.ToLower(str)
}

func lrnStrTrim() string {
	str := "NHello,how are you?NNNN" //to remove the specific char from the start or end ofstr, we can use TrimSpace to remove the unwanted spaces
	return strings.Trim(str, "N")
}

func main() {
	fmt.Printf("%v\n", learningStrContain())
	fmt.Printf("%v\n", lrnStrHassuffix())
	fmt.Printf("%v\n", lrnStrLower())
	fmt.Printf("%v\n", lrnStrTrim())
}
