package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}
