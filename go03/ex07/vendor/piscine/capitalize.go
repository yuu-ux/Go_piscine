package piscine

func IsLower(c rune) bool {
        return ('a' <= c && 'z' >= c)
}

func IsUpper(c rune) bool {
    return ('A' <= c && 'Z' >= c)
}

func IsDigit(c rune) bool {
	return ('0' <= c && '9' >= c)
}

func IsUpperLowerDigit (c rune) bool  {
	return (IsUpper(c) || IsLower(c) || IsDigit(c))
}

func Capitalize(s string) string {
	var result string
	var temp rune = '0'

	for i, c := range s {
		if i == 0 && IsLower(c) {
			c -= ('a' - 'A')
		}
		if IsUpper(c) && i != 0 {
			c += ('a' - 'A')
		}
		if !IsUpperLowerDigit(temp) && IsLower(c) {
			c -= ('a' - 'A')
		}
		result += string(c)
		temp = c
	}
	return result
}
