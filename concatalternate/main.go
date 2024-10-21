package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
	}

	for i, v := range slice1 {
		result = append(result, v)

		if i < len(slice2) {
			result = append(result, slice2[i])
		}
	}
	return result
}
func main() {
	args := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{1, 2, 3}, {4, 5}},
		{{}, {4, 5, 6}},
		{{1, 2, 3}, {}},
		{{}, {}},
		{{1, 2, 4}, {10, 20, 30, 40, 50}},
		{{2, 4, 6, 8, 10}, {1, 3, 5, 7, 9, 11}},
	}
	for _, arg := range args {
		fmt.Println("ConcatAlternate", ConcatAlternate(arg[0], arg[1]))
	}
}
