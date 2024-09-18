package main

import "fmt"

func PrintComb() {
	for p := '0'; p <= '9'; p++ {
		for o := '0'; o <= '9'; o++ {
			for w := '0'; w <= '9'; w++ {
				if p < o && o < w {
					fmt.Print(p)
					fmt.Print(o)
					fmt.Print(w)
					if p == '7' && o == '8' && w == '9' {
						fmt.Print('\n')
					} else {
						fmt.Print(',')
						fmt.Print(' ')
					}
				}
			}
		}
	}
}
