package piscine 

func BasicAtoi2(s string) int {
	var result int = 0
	for _, s := range s {
		if s >= '0' && s <= '9' {
			result = (result * 10) + (int(s) - '0')
		} else {
			return 0
		}
	}
	return result
}
