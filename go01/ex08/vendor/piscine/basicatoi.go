package piscine 

func BasicAtoi(s string) int {
	var result int = 0
	for _, s := range s {
		result = (result * 10) + (int(s) - '0')
	}
	return result
}