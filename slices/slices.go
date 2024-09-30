package main

import "fmt"

func main() {
	vabu := []string{"abacdafxi", "gafakjfe", "ihgighaig", "giafhaif", "aifhqi"}
	vabu = append(vabu, "aifhqi")
	vabu = append(vabu[:1], vabu[2:]...)
	fmt.Println(vabu)

	check := false
	for _, choco := range vabu {
		if choco == "aifhqi" {
			check = true
			break
		}
	}
	fmt.Println("check str:", check)
}
