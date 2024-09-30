package main

import "fmt"

func PrintIfNot(str string) string {
	if len(str) <= 3 {
		return "G\n"
	} else {
		return "Invalid Input  \n"
	}
}

func main() {
	table := []string{"ab", "		", "good morning sunshine!", "  28", "he", "", "honey!", "Z", "email123@live.fr", "w45m$", "-552", "474abc", "<="}
	for _, s := range table {
		fmt.Print(s + ":" + PrintIfNot(s))
	}
}
