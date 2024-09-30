package main

import "fmt"

func RectPerimeter(w, h int) int {
	if w <= 0 || h <= 0 {
		return -1
	}
	return 2 * (w + h)
}
func main() {
	args := [][]int{
		{1, 2},
		{0, 0},
		{8423, 9485},
		{-1, -1},
		{2147483647, 3},
		{14342, -1},
	}
	for _, arg := range args {
		fmt.Println(RectPerimeter(arg[0], arg[1]))
	}
}
