package piscine

import "ft"

func CountDigit(base string) int {
	var digit int
	for range base {
		digit++
	}
	return digit
}

func CheckBase(base []rune) bool {
	var cnt int
	for i, c1 := range base {
		if c1 == '+' || c1 == '-' {
			return false
		}
		for _, c2 := range base[i+1:] {
			if c1 == c2 {
				return false
			}
		}
		cnt++
	}
	if cnt < 2 {
		return false
	}
	return true
}

func PrintIndex(base []rune, index int) {
	ft.PrintRune(base[index])
}

func PutNbr(nbr int, digit int, base []rune) {
	if nbr != 0 {
		PutNbr(nbr / digit, digit, base)
		PrintIndex(base, nbr % digit)
	}
}

func PrintNbrBase(nbr int, base string) {
	var digit = CountDigit(base)

	if !CheckBase([]rune(base)) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	if nbr < 0 {
		nbr *= -1
		ft.PrintRune('-')
	}
	PutNbr(nbr, digit, []rune(base))
}
