package piscine

func IsLower(c rune) bool {
	return ('a' <= c && 'z' >= c)
}

func ToUpper(s string) string {
	var result string

	for _, c := range s {
		if IsLower(c) {
			c -= ('a' - 'A')
		}
		result += string(c)
	}
	return result
}
