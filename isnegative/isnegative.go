package main

import "fmt"

func IsNegative(nb int) {
	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}
func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}

/*func main() {
	table := append(
		random.IntSlice(),
		random.MinInt,
		random.MaxInt,
		0,
	)
	for _, arg := range table {
		challenge.Function("IsNegative", student.IsNegative, solutions.IsNegative, arg)
	}
}*/
