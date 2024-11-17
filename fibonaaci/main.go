package main

import "fmt"

func Fibonnaci(input int) []int {
	if input <= 0 {
		return []int{}
	}
	if input == 1 {
		return []int{0}
	}
	series := []int{0, 1}
	for i := 2; i < input; i++ {
		series = append(series, series[i-1]+series[i-2])
	}
	return series
}
func main() {
	n := 20
	fmt.Println(Fibonnaci(n))
}
