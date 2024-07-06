package piscine

func CountDigit(base string) int {
	var digit int
	for range base {
		digit++
	}
	return digit
}

func CheckBase(base []rune) bool {
	for i, c1 := range base {
		if c1 == '+' || c1 == '-' {
			return false
		}
		for _, c2 := range base[i+1:] {
			if c1 == c2 {
				return false
			}
		}
	}
	return true
}

func SearchIndex (c rune, base string) int {
	for i, c1 := range base {
		if c == c1 {
			return i
		}
	}
	return 0
}

func StringToInt(c rune) int {
	return (int(c) - '0')
}

func AtoiBase(s string, base string) int {
	var digit = CountDigit(base)
	var result int
	var index int

	if !(CheckBase([]rune(base))) || (digit < 2) {
		return 0
	}

	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			index = SearchIndex(c, base)
		} else {
			index = StringToInt(c)
		}
		result = (result * digit) + index
	}
	return result
}
