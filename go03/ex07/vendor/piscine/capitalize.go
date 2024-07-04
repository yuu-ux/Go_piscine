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
	return (IsUpper && IsLower && IsDigit)
}

func Capitalize(s string) string {
	var result string
	flag := true
	for _, c := range s {
		if flag == true {
			flag = false
			c -= ('a' - 'A')
		}
		if !IsUpperLowerDigit(c) {
			c -= ('a' - 'A')		
		}
		result += string(c)
	}
	
}
