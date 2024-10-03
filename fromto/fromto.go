package main

import (
	"fmt"
	"strconv"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "INVALID" + "\n"
	}
	result := ""
	if from == to {
		if from < 10 {
			result += "0" + strconv.Itoa(from)
		} else {
			result += strconv.Itoa(from)
		}
	}
	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + strconv.Itoa(i)
			} else {
				result += strconv.Itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	}

	if from > to {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0" + strconv.Itoa(i)
			} else {
				result += strconv.Itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	}
	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
	fmt.Print(FromTo(29, 17))
	fmt.Print(FromTo(99, 69))
}