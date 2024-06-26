package piscine

import "ft"

func PointOne(nb *int) {
	*nb = 1
	ft.PrintRune(rune(*nb + '0'))
}
