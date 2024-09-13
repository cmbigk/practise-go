package main

func StrLen(s string) int {
	counter := 0
	for _ = range s {
		counter++
	}
	return counter
}
