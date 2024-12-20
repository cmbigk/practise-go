package main

import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}

	if n%2 == 0 {
		return "fish"
	}

	if n%3 == 0 {
		return "chips"
	} else {
		return "error: non divisible"
	}

}

func main() {
	fmt.Println(FishAndChips(15))
	fmt.Println(FishAndChips(7))
	fmt.Println(FishAndChips(-5))
}
