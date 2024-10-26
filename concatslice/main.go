package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}

func ConcatSlice(slice1, slice2 []int) []int {
	var result []int

	for j := 0; j < len(slice1); j++ {
		result = append(result, slice1[j])
	}

	for i := 0; i < len(slice2); i++ {
		result = append(result, slice2[i])
	}

	return result

}
