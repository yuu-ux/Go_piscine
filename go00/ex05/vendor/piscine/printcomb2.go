package piscine

import "ft"

func PrintComb2() {
	for i:= 0; i < 10; i++ {
		for k := i + 1; k < 10; k++ {
			ft.PrintRune(rune(i + '0'))
			ft.PrintRune(rune(k + '0'))
			if i == 8 {
				break
			} else {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}
