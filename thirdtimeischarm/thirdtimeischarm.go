package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "" + "\n"

	}
	var result string
	for i := 2; i < len(str); i += 3 { //starts at index 2 bcuz we want to start from third character of the string, then just add 3 in that!!
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
