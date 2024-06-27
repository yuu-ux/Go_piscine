package piscine 

func Atoi(s string) int {
	var result int = 0
	sign := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, s := range s {
		if s >= '0' && s <= '9' {
			result = (result * 10) + (int(s) - '0')
		} else {
			return 0
		}
	}
	return result * sign
}