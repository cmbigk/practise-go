package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for p := '0'; p <= '9'; p++ {
		for o := '0'; o <= '9'; o++ {
			for w := '0'; w <= '9'; w++ {
				for e := '0'; e <= '9'; e++ {
					if p == '9' && o == '8' && w == '9' && e == '9' {
						z01.PrintRune(p)
						z01.PrintRune(o)
						z01.PrintRune(' ')
						z01.PrintRune(w)
						z01.PrintRune(e)
						z01.PrintRune('\n')
						break
					}
					if p == w && o < e || p < w {
						z01.PrintRune(p)
						z01.PrintRune(o)
						z01.PrintRune(' ')
						z01.PrintRune(w)
						z01.PrintRune(e)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
