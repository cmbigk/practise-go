package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return "" + "\n"
	}
	if len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i = i + 3 {
		result += string(str[i])
	}
	return result + "\n"
}
func main() {
	table := []string{"1234556789", "QKplq%QSw", "", "Kimetsu no Yaiba", "Z", "email123@live.fr", "8595485-52", "-552", "abc", "w58tw7474abc", "fifa world cup `2022`", "a b c d e f"}

	for _, s := range table {
		fmt.Println(ThirdTimeIsACharm(s))
	}
}
