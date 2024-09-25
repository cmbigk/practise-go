package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]
	if len(old) != 1 || len(new) != 1 {
		return
	}
	if strings.Contains(str, old) {
		str = strings.ReplaceAll(str, old, new)
	}
	fmt.Println(str)

}
