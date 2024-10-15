package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	var result []int
	// Concatenate the two slices
	result = append(result, slice1...)
	result = append(result, slice2...)
	// Return the concatenated slice
	return result

}

func main() {
	args := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{}, {-10, 0, 2}},
		{{-10, 0, 2}, {}},
		{{}, {}},
		{{1, 2, 3}, {4, 5, 6, 3, 4, 5, 6}},
		{{0, 0, 0}, {0, 0, 0}},
	}

	for _, arg := range args {
		fmt.Println("ConcatSlice", ConcatSlice(arg[0], arg[1]))
	}
}
