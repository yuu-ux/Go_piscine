package piscine

import "ft"

func PrintComb() {
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				ft.PrintRune(rune(i + '0'))
				ft.PrintRune(rune(j + '0'))
				ft.PrintRune(rune(k + '0'))
				if i == 7 {
					break	
				} else {
					ft.PrintRune(44)
					ft.PrintRune(32)
				}
			}
		}
	}
	ft.PrintRune('\n')
}
