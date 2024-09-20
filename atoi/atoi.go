package main

import "fmt"

func Atoi(s string) int {
	sign := 1
	if len(s) == 0 {
		return 0
	}
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}

	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		}
		result = (result*10 + int(char-'0'))
	}
	return result * sign
}
func main() {
	table := []string{"",
		"-",
		"+",
		"0",
		"+0",
		"-Invalid123",
		"--123",
		"-+123",
		"+657",
		"++123",
		"123-",
		"123+",
		"123.",
		"123.0",
		"123a45",
		"+1234",
		"-1234",
		"+123456",
		"-123456",
	}

	for _, arg := range table {
		fmt.Println(Atoi(arg))
	}
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))

}
