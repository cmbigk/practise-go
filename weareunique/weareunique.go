package main

import "fmt"

func isContain(g string, d rune) bool {
	for _, b := range g {
		if b == d {
			return true
		}
	}
	return false
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	if str1 == str2 {
		return 0
	}

	result := ""
	for _, v := range str1 {
		if !isContain(str2, v) && !isContain(result, v) {
			result += string(v)
		}

		for _, r := range str2 {
			if !isContain(str1, r) && !isContain(result, r) {
				result += string(r)
			}
		}
	}
	return len(result)
}

func main() {
	table := [][]string{
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"abc", "abc"},
		{"", ""},
		{"pomme", "pomme"},
		{"+265", "265"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
		{"", "eveRyone"},
		{"_55w1se", "55w1se"},
		{"foo", "boo"},
	}
	for _, arg := range table {
		fmt.Println("Your arg is:", arg, WeAreUnique(arg[0], arg[1]))
	}
}
