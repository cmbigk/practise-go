package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	var result []int
	for _, v := range slice1 {
		result = append(result, v)
	}

	for _, v := range slice2 {
		result = append(result, v)
	}

	//
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
		{{1, 2, 3}, {4, 5, 6}},
	}

	for _, arg := range args {
		fmt.Println("ConcatSlice", ConcatSlice(arg[0], arg[1]))
	}
}
