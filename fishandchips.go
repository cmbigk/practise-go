package main

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	if (n%2 == 0) && (n%3 == 0) {
		return "Fish and Chips"
	}
	if n%2 == 0 {
		return "Fish"
	}
	if n%3 == 0 {
		return "Chips"
	}
	return "error: non divisible"
}
