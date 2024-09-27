package main

import "fmt"

func Iscapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' && s[i] != 0 && s[i-1] == ' ' {
			return false
		}
	}
	return !(s[0] >= 'a' && s[0] <= 'z')
}
func main() {
	fmt.Println(Iscapitalized("Hello! How are you?"))
	fmt.Println(Iscapitalized("Hello How Are You"))
	fmt.Println(Iscapitalized("Whats 4this 100K?"))
	fmt.Println(Iscapitalized("Whatsthis4"))
	fmt.Println(Iscapitalized("!!!!Whatsthis4"))
	fmt.Println(Iscapitalized(""))
}
