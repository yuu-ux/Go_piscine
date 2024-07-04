package piscine

func IsDigit(c rune) bool {
	return ('0' <= c && '9' >= c)
}

func IsNumeric(s string) bool {
    for _, c := range s {
        if !IsDigit(c) {
            return false
        }
    }
    return true
}
