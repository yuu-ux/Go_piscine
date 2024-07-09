package piscine

func CountDigit(base string) int {
    var digit int
    for range base {
        digit++
    }
    return digit
}

func ReturnNbr(nbr int, digit int, base []rune, result string) string {
    if nbr != 0 {
        result = ReturnNbr(nbr / digit, digit, base, result)
        result += string(base[nbr % digit])
    }
	return result
}

func SearchIndex (c rune, base string) int {
    for i, c1 := range base {
        if c == c1 {
            return i
        }
    }
    return 0
}

func AtoiBase(s string, base string) int {
    var digit = CountDigit(base)
    var result int
    var index int

    for _, c := range s {
        if !(c >= '0' && c <= '9') {
            index = SearchIndex(c, base)
        } else {
            index = (int(c) - '0')
		}
        result = (result * digit) + index
    }
    return result
}

func PrintNbrBase(nbr int, base string) string {
    var digit = CountDigit(base)
	var result string

    if nbr < 0 {
        nbr *= -1
    }
    return ReturnNbr(nbr, digit, []rune(base), result)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := AtoiBase(nbr, baseFrom)
	result := PrintNbrBase(decimal, baseTo)
	return result
}
