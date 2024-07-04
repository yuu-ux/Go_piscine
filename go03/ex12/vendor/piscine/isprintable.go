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
func IsPrintable(s string) bool {
    for _, c := range s {
        if !('a' <= c && 'z' >= c) {
            return false
        }
    }
    return true
}
