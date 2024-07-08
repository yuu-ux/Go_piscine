package piscine

func SearchSeparators(c rune) bool {
	if c == ' ' || c == '	' ||  c == '\n' {
		return true
	}
	return false
}

func SplitWhiteSpaces(s string) []string {
	var temp string
	var result []string

	for _, c := range s {
		if SearchSeparators(c) {
			result = append(result, temp)
			temp = ""
			continue
		}
		temp += string(c)
	}
	result = append(result, temp)
	return result
}
