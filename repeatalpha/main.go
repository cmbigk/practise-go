package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'a' && c < +'z' {
			for i := 0; i < int((c-'a')+1); i++ {
				result += string(c)
			}
		} else if c >= 'A' && c <= 'Z' {
			for i := 0; i < int((c-'A')+1); i++ {
				result += string(c)
			}
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	testCases := []struct {
		in   string
		want string
	}{
		{"abc", "abbccc"},
		{"Choumi.", "CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii."},
		{"", ""},
		{"abacadaba 01!", "abbacccaddddabba 01!"},
	}
	for _, tc := range testCases {
		got := RepeatAlpha(tc.in)
		fmt.Printf("RepeatAlpha(%q) = %q, expected %q\n", tc.in, got, tc.want)
	}
}
