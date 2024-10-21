package main

import "fmt"

func RepeatAlpha(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			repeatchar := int(c-'a') + 1
			for i := 0; i < repeatchar; i++ {
				result += string(c)
			}
		} else if c >= 'A' && c <= 'Z' {
			repeatchar := int(c-'A') + 1
			for i := 0; i < repeatchar; i++ {
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
