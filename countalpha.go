package main

func CountAlpha(s string) int {
	count := 0
	for _, k := range s {
		if (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') {
			count++
		}
	}
	return count
}
