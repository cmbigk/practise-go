package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			for i := 0; i < int(r-'a'+1); i++ {
				result += string(r)
			}
		} else if r >= 'A' && r <= 'Z' {
			for i := 0; i < int(r-'A'+1); i++ {
				result += string(r)
			}
		} else {
			result += string(r)
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
