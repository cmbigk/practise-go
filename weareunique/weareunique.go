package main

import "fmt"

func isContain(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
func WeAreUnique(str1, str2 string) (int, string) {
	if str1 == "" && str2 == "" {
		return -1, ""
	}
	result := ""
	for _, v := range str1 {
		if !isContain(str2, v) && !isContain(result, v) {
			result += string(v)
		}
	}

	for _, v := range str2 {
		if !isContain(str1, v) && !isContain(result, v) {
			result += string(v)
		}
	}
	return len(result), ""
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
		fmt.Println(WeAreUnique(arg[0], arg[1]))
		fmt.Println(WeAreUnique(arg[0], arg[1]))
		fmt.Println(WeAreUnique(arg[0], arg[1]))
	}
}
