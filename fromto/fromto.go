package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
	result := ""
	if from < 0 || to < 0 || from > 99 || to > 99 {
		return "Invalid"
	}
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
