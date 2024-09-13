package main

func CountChar(str string, c rune) int {
	counter := 0
	for _, k := range str {
		if k == c {
			counter++
		}
	}
	return counter
}
